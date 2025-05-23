package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/solumD/go-quotes-server/internal/config"
	"github.com/solumD/go-quotes-server/internal/handler"
	"github.com/solumD/go-quotes-server/internal/lib/middleware"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/repository/postgres"
	"github.com/solumD/go-quotes-server/internal/service/srv"
)

const (
	serverReadTimeOut  = 5 * time.Second
	serverWriteTimeOut = 5 * time.Second
	serverIdleTimeOut  = 60 * time.Second
)

// InitAndRun initializes the application and runs it.
func InitAndRun(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg := config.MustLoad()

	logger := sl.InitLogger(cfg.LoggerLevel())
	logger.Info("starting quotes server")
	logger.Debug("debug messages are enabled")

	repository, err := postgres.New(ctx, cfg.PostgresDSN())
	if err != nil {
		logger.Error("failed to create repository", sl.Err(err))
		os.Exit(1)
	}
	defer repository.Close()
	logger.Info("connected to database")

	service := srv.New(repository)
	logger.Info("service initialized")

	handler := handler.New(service)
	logger.Info("handler initialized")

	router := mux.NewRouter()
	router.Use(middleware.NewMWLogger(logger))
	logger.Info("logger middleware enabled")

	router.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		if author != "" {
			handler.GetQuotesByAuthor(ctx, logger)(w, r)
		} else {
			handler.GetAllQuotes(ctx, logger)(w, r)
		}
	}).Methods("GET")
	router.HandleFunc("/quotes", handler.SaveQuote(ctx, logger)).Methods("POST")
	router.HandleFunc("/quotes/random", handler.GetRandomQuote(ctx, logger)).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuote(ctx, logger)).Methods("DELETE")
	logger.Info("initialized routes")

	server := &http.Server{
		Addr:         cfg.ServerAddr(),
		Handler:      router,
		ReadTimeout:  serverReadTimeOut,
		WriteTimeout: serverWriteTimeOut,
		IdleTimeout:  serverIdleTimeOut,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("failed to start server", sl.Err(err))
		}
	}()

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Shutdown(ctx)
	logger.Info("shutting down")
	os.Exit(0)
}
