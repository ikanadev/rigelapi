package extra

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(app *fiber.App, teacherApp fiber.Router, ent *ent.Client, ctx context.Context) {
	repo := NewExtraEntRepo(ent, ctx)
	handlers := NewExtraHandler(app, teacherApp, repo)
	handlers.handle()
}
