package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/solumD/go-quotes-server/internal/config"
	"github.com/solumD/go-quotes-server/internal/handler"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/repository/postgres"
	"github.com/solumD/go-quotes-server/internal/service/srv"
	"golang.org/x/sync/errgroup"
)

const (
	serverReadTimeOut  = 5 * time.Second
	serverWriteTimeOut = 5 * time.Second
	serverIdleTimeOut  = 60 * time.Second
)

func InitAndRun() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

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
	router.PathPrefix("/quotes")
	router.HandleFunc("/", handler.SaveQuote(ctx, logger)).Methods("POST")
	router.HandleFunc("/", handler.GetAllQuotes(ctx, logger)).Methods("GET")
	router.HandleFunc("/random", handler.GetRandomQuote(ctx, logger)).Methods("GET")
	router.HandleFunc("/{author}", handler.GetQuotesByAuthor(ctx, logger)).Methods("GET")
	router.HandleFunc("/", handler.DeleteQuote(ctx, logger)).Methods("DELETE")
	logger.Info("initialized routes")

	server := &http.Server{
		Addr:         cfg.ServerAddr(),
		Handler:      router,
		ReadTimeout:  serverReadTimeOut,
		WriteTimeout: serverWriteTimeOut,
		IdleTimeout:  serverIdleTimeOut,
	}

	// server start and graceful shutdown
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		logger.Info("starting server", slog.String("addr", cfg.ServerAddr()))
		return server.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		logger.Info("os exit call", sl.Err(err))
	}

}
