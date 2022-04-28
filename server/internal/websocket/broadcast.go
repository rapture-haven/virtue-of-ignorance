package websocket

import (
	"rapture-haven/virtue-of-ignorance/internal/types"
	"rapture-haven/virtue-of-ignorance/internal/types/event"
	"rapture-haven/virtue-of-ignorance/internal/types/message"
)

func (h *MessageHandler) boradcastCreate(session types.Session, payload *message.BroadcastCreate) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	pipelineErr := h.broadcast.Create(payload.URL)
	if pipelineErr != nil {
		if err := session.Send(
			message.SystemMessage{
				Event:   event.SYSTEM_ERROR,
				Title:   "Error while starting broadcast",
				Message: pipelineErr.Error(),
			}); err != nil {
			h.logger.Warn().Err(err).Msgf("sending event %s has failed", event.SYSTEM_ERROR)
			return err
		}
	}

	if err := h.boradcastStatus(nil); err != nil {
		return err
	}

	return nil
}

func (h *MessageHandler) boradcastDestroy(session types.Session) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	h.broadcast.Destroy()

	if err := h.boradcastStatus(nil); err != nil {
		return err
	}

	return nil
}

func (h *MessageHandler) boradcastStatus(session types.Session) error {
	// if no session, broadcast change
	if session == nil {
		if err := h.sessions.AdminBroadcast(
			message.BroadcastStatus{
				Event:    event.BORADCAST_STATUS,
				IsActive: h.broadcast.IsActive(),
				URL:      h.broadcast.GetUrl(),
			}, nil); err != nil {
			h.logger.Warn().Err(err).Msgf("broadcasting event %s has failed", event.BORADCAST_STATUS)
			return err
		}

		return nil
	}

	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	if err := session.Send(
		message.BroadcastStatus{
			Event:    event.BORADCAST_STATUS,
			IsActive: h.broadcast.IsActive(),
			URL:      h.broadcast.GetUrl(),
		}); err != nil {
		h.logger.Warn().Err(err).Msgf("sending event %s has failed", event.BORADCAST_STATUS)
		return err
	}

	return nil
}
