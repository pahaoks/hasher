package app

import (
	"context"
	"testing"
	"time"
)

func TestStartStopContextCancel(t *testing.T) {
	app := New()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 1)

		cancel()
	}()

	app.Start(ctx)
}

func TestStartStop(t *testing.T) {
	app := New()

	go func() {
		time.Sleep(time.Second * 1)

		app.Stop()
	}()

	app.Start(context.Background())
}
