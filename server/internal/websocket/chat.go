package websocket

import (
	"rapture-haven/virtue-of-ignorance/internal/types"
	"rapture-haven/virtue-of-ignorance/internal/types/event"
	"rapture-haven/virtue-of-ignorance/internal/types/message"
)

func (h *MessageHandler) chat(id string, session types.Session, payload *message.ChatReceive) error {
	if session.Muted() {
		return nil
	}

	if err := h.sessions.Broadcast(
		message.ChatSend{
			Event:   event.CHAT_MESSAGE,
			Content: payload.Content,
			ID:      id,
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_RELEASE)
		return err
	}
	return nil
}

func (h *MessageHandler) chatEmote(id string, session types.Session, payload *message.EmoteReceive) error {
	if session.Muted() {
		return nil
	}

	if err := h.sessions.Broadcast(
		message.EmoteSend{
			Event: event.CHAT_EMOJI,
			Emote: payload.Emote,
			ID:    id,
		}, nil); err != nil {
		h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.CONTROL_RELEASE)
		return err
	}
	return nil
}
