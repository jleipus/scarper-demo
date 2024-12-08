package service

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func Run(run func(context.Context) error) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
