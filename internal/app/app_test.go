package app

import (
	"context"
	"testing"
	"time"
)

func TestStartStopApp(t *testing.T) {
	app := New()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 1)

		cancel()
	}()

	app.Start(ctx)
}
