package class

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(app *fiber.App, teacherApp fiber.Router, ent *ent.Client, ctx context.Context) {
	repo := NewClassEntRepo(ent, ctx)
	handlers := NewClassHandler(app, teacherApp, repo)
	handlers.handle()
}
