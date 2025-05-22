package main

import (
	"context"

	"github.com/solumD/go-quotes-server/internal/app"
)

func main() {
	ctx := context.Background()
	app.InitAndRun(ctx)
}
