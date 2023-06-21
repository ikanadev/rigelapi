package teacher

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func Start(app fiber.Router, ent *ent.Client, ctx context.Context) {
	repo := NewTeacherEntRepo(ent, ctx)
	handlers := NewTeacherHandler(app, repo)
	handlers.handle()
}
