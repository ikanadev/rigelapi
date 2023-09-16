package sync

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/common"
)

type SyncHandler struct {
	teacherApp fiber.Router
	repo       SyncRepository
}

func NewSyncHandler(teacherApp fiber.Router, repo SyncRepository) SyncHandler {
	return SyncHandler{teacherApp, repo}
}

func (sh *SyncHandler) handle() {
	sh.teacherApp.Post("/students", sh.handleSyncStudents)
	sh.teacherApp.Get("/students/year/:yearid", sh.handleGetStudents)
	sh.teacherApp.Get("/classperiods/year/:yearid", sh.handleGetClassPeriods)
}

func (sh *SyncHandler) handleSyncStudents(ctx *fiber.Ctx) error {
	var students []common.StudentTx
	err := ctx.BodyParser(&students)
	if err != nil {
		return err
	}
	err = sh.repo.SyncStudents(students)
	return err
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

func (sh *SyncHandler) handleGetClassPeriods(ctx *fiber.Ctx) error {
	teacherID, ok := ctx.Locals("id").(string)
	if !ok || len(teacherID) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	yearID := ctx.Params("yearid")
	classPeriods, err := sh.repo.GetClassPeriods(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(classPeriods)
}
