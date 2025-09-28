package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create context that gets cancelled on certain signals.
	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer cancelFn()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v", err)
		cancelFn()
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return nil
}
