package main

import (
	"context"

	"example.com/hasher/internal/app"
)

func main() {
	app.New().Start(context.Background())
}
