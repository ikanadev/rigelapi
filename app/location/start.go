package location

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(app *fiber.App, ent *ent.Client, ctx context.Context) {
	entRepo := NewLocationEntRepo(ent, ctx)
	handlers := NewLocationHandler(app, entRepo)
	handlers.handle()
}
