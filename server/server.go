package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alvinfebriando/costumer-test/config"
)

func New(router http.Handler) *http.Server {
	restConfig := config.NewRestConfig()
	addr := fmt.Sprintf("%s:%s", restConfig.Host, restConfig.Port)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartWithGracefulShutdown(s *http.Server) {
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	timeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
