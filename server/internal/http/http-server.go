package http

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rapture-haven/virtue-of-ignorance/internal/config"
	"github.com/rapture-haven/virtue-of-ignorance/internal/types"
)

type HTTPServer struct {
	logger zerolog.Logger
	router *chi.Mux
	http   *http.Server
	conf   *config.Server
}

func New(conf *config.Server, webSocketHandler types.WebSocketHandler) *HTTPServer {
	logger := log.With().Str("module", "http").Logger()

	router := chi.NewRouter()
	router.Use(middleware.RequestID) // Create a request ID for each request
	router.Use(middleware.RequestLogger(&logformatter{logger}))
	router.Use(middleware.Recoverer) // Recover from panics without crashing server

	router.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := webSocketHandler.Upgrade(w, r)
		if err != nil {
			logger.Warn().Err(err).Msg("failed to upgrade websocket conection")
		}
	})

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("true"))
	})

	fs := http.FileServer(http.Dir(conf.Static))
	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(conf.Static + r.URL.Path); !os.IsNotExist(err) {
			fs.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	server := &http.Server{
		Addr:    conf.Bind,
		Handler: router,
	}

	return &HTTPServer{
		logger: logger,
		router: router,
		http:   server,
		conf:   conf,
	}
}

func (s *HTTPServer) Start() {
	if s.conf.Cert != "" && s.conf.Key != "" {
		go func() {
			if err := s.http.ListenAndServeTLS(s.conf.Cert, s.conf.Key); err != http.ErrServerClosed {
				s.logger.Panic().Err(err).Msg("unable to start https server")
			}
		}()
		s.logger.Info().Msgf("https listening on %s", s.http.Addr)
	} else {
		go func() {
			if err := s.http.ListenAndServe(); err != http.ErrServerClosed {
				s.logger.Panic().Err(err).Msg("unable to start http server")
			}
		}()
		s.logger.Warn().Msgf("http listening on %s", s.http.Addr)
	}
}

func (s *HTTPServer) Shutdown() error {
	return s.http.Shutdown(context.Background())
}
