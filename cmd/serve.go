package cmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mdmoshiur/example-go/internal/cache"
	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/middleware"
	"github.com/mdmoshiur/example-go/internal/validation"
	"github.com/spf13/cobra"

	apphttp "github.com/mdmoshiur/example-go/http"
)

const serverShutDownContextDeadline = 10

// serveCmd represents the serve command
var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve run available servers such as: HTTP/JSON or gRPC",
		Long:  `Serve run available servers such as: HTTP/JSON or gRPC`,
		PreRun: func(cmd *cobra.Command, args []string) {
			logger.Info("Connecting database...")
			if err := conn.ConnectDB(); err != nil {
				logger.Fatal(err)
			}
			logger.Info("Hola, Database connected successfully!")

			logger.Info("Connecting redis...")
			if err := conn.ConnectRedis(); err != nil {
				logger.Fatal(err)
			}
			logger.Info("Redis connected successfully!")

			logger.Info("Initializing validator...")
			if err := validation.InitValidator(); err != nil {
				logger.Fatal(err)
			}
			logger.Info("Validator initialization done!")
		},
		Run: serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	// boot http server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// build and run http server
	srv := buildHTTP(cmd, args)
	go func(sr *http.Server) {
		if err := sr.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err)
		}
	}(srv)

	<-stop
	logger.Info("Shutting down http server...")
	ctx, cancel := context.WithTimeout(context.Background(), serverShutDownContextDeadline*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); !errors.Is(err, http.ErrServerClosed) {
		logger.Fatal(err)
	}
	logger.Info("Server shutdown successful!")
}

// buildHTTP register available handlers and return a http server
func buildHTTP(cmd *cobra.Command, args []string) *http.Server {
	r := chi.NewRouter()

	// cors middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// middlewares
	r.Use(
		chimiddleware.Logger,
		chimiddleware.Recoverer,
		chimiddleware.Heartbeat("/"),
	)

	cfg := config.App()
	db := conn.DefaultDB()
	defaultCache := conn.DefaultRedis()
	cs := cache.NewRedisCache(defaultCache, config.Redis().Prefix, config.Redis().DefaultTTL)

	r.Use(
		middleware.WithDB(db),
		middleware.WithCache(cs),
	)

	handler := apphttp.RegisterHandlers(db, cs)
	apphttp.RegisterRoutes(r, handler)

	logger.Info("HTTP listening on port: ", cfg.HTTPPort)
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:           r,
		ReadHeaderTimeout: cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}
