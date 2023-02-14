package servers

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Wait for interruption events, then gracefully execute callback
func onShutdown(ctx context.Context, callback func()) {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case <-s:
	case <-ctx.Done():
	}

	callback()
}
