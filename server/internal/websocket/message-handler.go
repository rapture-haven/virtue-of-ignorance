package websocket

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/rapture-haven/virtue-of-ignorance/internal/types"
	event "github.com/rapture-haven/virtue-of-ignorance/internal/types/events"
	message "github.com/rapture-haven/virtue-of-ignorance/internal/types/messages"
)

type MessageHandler struct {
	logger         zerolog.Logger
	sessionManager types.SessionManager
	locked         map[string]string
}

func (h *MessageHandler) Connected(admin bool, socket *WebSocket) (bool, string) {
	address := socket.Address()
	if address == "" {
		h.logger.Debug().Msg("no remote address")
	}

	_, ok := h.locked["login"]
	if ok && !admin {
		h.logger.Debug().Msg("server locked")
		return false, "locked"
	}

	return true, ""
}

func (h *MessageHandler) Disconnected(id string) {
	h.sessionManager.Destroy(id)
}

func (h *MessageHandler) Message(id string, raw []byte) error {
	header := message.Message{}
	if err := json.Unmarshal(raw, &header); err != nil {
		return err
	}

	_, ok := h.sessionManager.Get(id)
	if !ok {
		return errors.Errorf("unknown session id %s", id)
	}

	// TODO: заглушка
	switch header.Event {
	default:
		return errors.Errorf("websocket api got unknown message event %s", header.Event)
	}
}

func (h *MessageHandler) SessionCreated(id string, session types.Session) error {

	// TODO: здесь будет инициализация webrtc пира + отправка sdp и id сессии клиенту

	// кидаем ивент о завершении инициализации
	if err := session.Send(message.SystemInit{
		Event: event.SYSTEM_INIT,
		Locks: h.locked,
	}); err != nil {
		h.logger.Warn().Str("id", id).Err(err).Msgf("sending event %s has failed", event.SYSTEM_INIT)
		return err
	}

	return nil
}

func (h *MessageHandler) SessionConnected(id string, session types.Session) error {

	// send list of members to session
	if err := session.Send(message.MembersList{
		Event:    event.MEMBER_LIST,
		Memebers: h.sessionManager.Members(),
	}); err != nil {
		h.logger.Warn().Str("id", id).Err(err).Msgf("sending event %s has failed", event.MEMBER_LIST)
		return err
	}

	return nil
}

func (h *MessageHandler) SessionDestroyed(id string) error {
	if h.sessionManager.IsHost(id) {
		h.sessionManager.ClearHost()
	}

	// TODO: нам нужен бродкаст ивента о том, что сессия уничтожена - всем участникам сессии

	return nil
}
