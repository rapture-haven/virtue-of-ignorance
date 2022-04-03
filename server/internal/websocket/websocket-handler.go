package websocket

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rapture-haven/virtue-of-ignorance/internal/config"
	"github.com/rapture-haven/virtue-of-ignorance/internal/types"
	event "github.com/rapture-haven/virtue-of-ignorance/internal/types/events"
	message "github.com/rapture-haven/virtue-of-ignorance/internal/types/messages"
	"github.com/rapture-haven/virtue-of-ignorance/internal/utils"
)

const CONTROL_PROTECTION_SESSION = "by_control_protection"

// Send pings to peer with this period. Must be less than pongWait.
const pingPeriod = 60 * time.Second

type WebSocketHandler struct {
	config   *config.WebSocket
	logger   zerolog.Logger
	wg       sync.WaitGroup
	shutdown chan interface{}

	// TODO: stats в зачаточном состоянии
	connsCnt uint32

	upgrader       websocket.Upgrader
	sessionManager types.SessionManager
	messageHandler *MessageHandler
}

func New(sessionManager types.SessionManager, config *config.WebSocket) *WebSocketHandler {

	logger := log.With().Str("module", "websocket").Logger()
	locks := make(map[string]string)

	return &WebSocketHandler{
		logger:         logger,
		shutdown:       make(chan interface{}),
		config:         config,
		sessionManager: sessionManager,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		messageHandler: &MessageHandler{
			logger:         logger.With().Str("subsystem", "messageHandler").Logger(),
			sessionManager: sessionManager,
			locked:         locks,
		},
	}
}

func (ws *WebSocketHandler) Start() {
	ws.sessionManager.OnCreated(func(id string, session types.Session) {
		if err := ws.messageHandler.SessionCreated(id, session); err != nil {
			ws.logger.Warn().Str("id", id).Err(err).Msg("session created with and error")
		} else {
			ws.logger.Debug().Str("id", id).Msg("session created")
		}
	})

	ws.sessionManager.OnConnected(func(id string, session types.Session) {
		if err := ws.messageHandler.SessionConnected(id, session); err != nil {
			ws.logger.Warn().Str("id", id).Err(err).Msg("session connected with and error")
		} else {
			ws.logger.Debug().Str("id", id).Msg("session connected")
		}
	})

	ws.sessionManager.OnDestroy(func(id string, session types.Session) {
		if err := ws.messageHandler.SessionDestroyed(id); err != nil {
			ws.logger.Warn().Str("id", id).Err(err).Msg("session destroyed with and error")
		} else {
			ws.logger.Debug().Str("id", id).Msg("session destroyed")
		}
	})

	// TODO: заглушка - вебсокет просто слушает, ничего не делает
	ws.wg.Add(1)
	go func() {
		defer func() {
			ws.logger.Info().Msg("shutdown")
			ws.wg.Done()
		}()

		for {
			select {
			case <-ws.shutdown:
				return
			default:
				time.Sleep(100 * time.Millisecond)

				if !ws.sessionManager.HasHost() {
					continue
				}
			}
		}
	}()
}

func (ws *WebSocketHandler) Shutdown() error {
	close(ws.shutdown)
	ws.wg.Wait()
	return nil
}

func (ws *WebSocketHandler) Upgrade(w http.ResponseWriter, r *http.Request) error {
	ws.logger.Debug().Msg("attempting to upgrade connection")

	id, err := utils.NewUID(32)
	if err != nil {
		ws.logger.Error().Err(err).Msg("failed to generate user id")
		return err
	}

	connection, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.logger.Error().Err(err).Msg("failed to upgrade connection")
		return err
	}

	admin, err := ws.authenticate(r)
	if err != nil {
		ws.logger.Warn().Err(err).Msg("authentication failed")

		if err = connection.WriteJSON(message.SystemMessage{
			Event:   event.SYSTEM_DISCONNECT,
			Message: "invalid_password",
		}); err != nil {
			ws.logger.Error().Err(err).Msg("failed to send disconnect")
		}

		if err = connection.Close(); err != nil {
			return err
		}
		return nil
	}

	socket := &WebSocket{
		id:         id,
		ws:         ws,
		address:    utils.GetHttpRequestIP(r),
		connection: connection,
	}

	ok, reason := ws.messageHandler.Connected(admin, socket)
	if !ok {
		if err = connection.WriteJSON(message.SystemMessage{
			Event:   event.SYSTEM_DISCONNECT,
			Message: reason,
		}); err != nil {
			ws.logger.Error().Err(err).Msg("failed to send disconnect")
		}

		if err = connection.Close(); err != nil {
			return err
		}

		return nil
	}

	ws.sessionManager.New(id, admin, socket)

	ws.logger.
		Debug().
		Str("session", id).
		Str("address", connection.RemoteAddr().String()).
		Msg("new connection created")

	atomic.AddUint32(&ws.connsCnt, uint32(1))

	defer func() {
		ws.logger.
			Debug().
			Str("session", id).
			Str("address", connection.RemoteAddr().String()).
			Msg("session ended")

		atomic.AddUint32(&ws.connsCnt, ^uint32(0))
	}()

	ws.handle(connection, id)
	return nil
}

func (ws *WebSocketHandler) IsAdmin(password string) (bool, error) {
	if password == ws.config.AdminPassword {
		return true, nil
	}

	if password == ws.config.Password {
		return false, nil
	}

	return false, fmt.Errorf("invalid password")
}

func (ws *WebSocketHandler) authenticate(r *http.Request) (bool, error) {
	passwords, ok := r.URL.Query()["password"]
	if !ok || len(passwords[0]) < 1 {
		return false, fmt.Errorf("no password provided")
	}

	return ws.IsAdmin(passwords[0])
}

func (ws *WebSocketHandler) handle(connection *websocket.Conn, id string) {
	bytes := make(chan []byte)
	cancel := make(chan struct{})
	ticker := time.NewTicker(pingPeriod)

	ws.wg.Add(1)
	go func() {
		defer func() {
			ticker.Stop()
			ws.logger.Debug().Str("address", connection.RemoteAddr().String()).Msg("handle socket ending")
			ws.messageHandler.Disconnected(id)
			ws.wg.Done()
		}()

		for {
			_, raw, err := connection.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					ws.logger.Warn().Err(err).Msg("read message error")
				} else {
					ws.logger.Debug().Err(err).Msg("read message error")
				}
				close(cancel)
				break
			}
			bytes <- raw
		}
	}()

	for {
		select {
		case raw := <-bytes:
			ws.logger.Debug().
				Str("session", id).
				Str("address", connection.RemoteAddr().String()).
				Str("raw", string(raw)).
				Msg("received message from client")
			if err := ws.messageHandler.Message(id, raw); err != nil {
				ws.logger.Error().Err(err).Msg("message messageHandler has failed")
			}
		case <-ws.shutdown:
			if err := connection.WriteJSON(message.SystemMessage{
				Event:   event.SYSTEM_DISCONNECT,
				Message: "server_shutdown",
			}); err != nil {
				ws.logger.Err(err).Msg("failed to send disconnect")
			}

			if err := connection.Close(); err != nil {
				ws.logger.Err(err).Msg("connection closed with an error")
			}
			return
		case <-cancel:
			return
		case <-ticker.C:
			if err := connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
