package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(
	app *fiber.App,
	ent *ent.Client,
	ctx context.Context,
	config config.Config,
	genID func() string,
) {
	entRepo := NewAuthEntRepo(ent, ctx, config, genID)
	handlers := NewAuthHandler(app, entRepo, config)
	handlers.handle()
}
