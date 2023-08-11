package main

import (
	"context"

	"github.com/think-free/sportsnews/lib/logging"

	"github.com/think-free/sportsnews/providers/htafc/internal/cliparams"
	"github.com/think-free/sportsnews/providers/htafc/internal/config"
	"github.com/think-free/sportsnews/providers/htafc/internal/database"
	"github.com/think-free/sportsnews/providers/htafc/internal/service"
	"github.com/think-free/sportsnews/providers/htafc/internal/upstream"
)

func main() {
	ctx := context.Background()
	c := config.New()
	cp := cliparams.New()

	logging.Init(cp.LogLevel)

	db := database.New(ctx, cp)
	up := upstream.New(ctx, cp, c)
	srv := service.New(ctx, c, up, db)
	srv.Run(ctx)
}
