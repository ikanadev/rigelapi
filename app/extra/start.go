package extra

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(app *fiber.App, ent *ent.Client, ctx context.Context) {
	repo := NewExtraEntRepo(ent, ctx)
	handlers := NewExtraHandler(app, repo)
	handlers.handle()
}
