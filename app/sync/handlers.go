package sync

import (
	"github.com/gofiber/fiber/v2"
)

type SyncHandler struct {
	teacherApp fiber.Router
	repo       SyncRepository
}

func NewSyncHandler(teacherApp fiber.Router, repo SyncRepository) SyncHandler {
	return SyncHandler{teacherApp, repo}
}

func (sh *SyncHandler) handle() {
	sh.teacherApp.Get("/students/year/:yearid", sh.handleGetStudents)
}

func (sh *SyncHandler) handleGetStudents(ctx *fiber.Ctx) error {
	teacherID, ok := ctx.Locals("id").(string)
	if !ok || len(teacherID) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	yearID := ctx.Params("yearid")
	students, err := sh.repo.GetStudents(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(students)
}
