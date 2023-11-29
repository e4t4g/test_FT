package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/e4t4g/test_FT/internal/config"
	"github.com/e4t4g/test_FT/internal/controller"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Errorf("logger error")
	}

	defer logger.Sync()

	sugar := logger.Sugar()

	cfg := config.NewConfig()

	srv := chi.NewRouter()

	r := &http.Server{Addr: fmt.Sprintf(":%s", cfg.Port), Handler: controller.NewController(srv, cfg, sugar)}

	go func(ctx context.Context, r *http.Server) {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = r.Shutdown(ctx); err != nil {
			sugar.Fatalf("Server forced to shutdown: %s", err)
		}
	}(ctx, r)

	if err = r.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		sugar.Fatalf("listen: %s\n", err.Error())
	}

	sugar.Info("Server exiting: ", ctx.Err())

}
