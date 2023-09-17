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
	sh.teacherApp.Get("/students/year/:yearid", sh.handleGetStudents)
	sh.teacherApp.Post("/students", sh.handleSyncStudents)
	sh.teacherApp.Get("/classperiods/year/:yearid", sh.handleGetClassPeriods)
	sh.teacherApp.Post("/classperiods", sh.handleSyncClassPeriods)
	sh.teacherApp.Get("/attendancedays/year/:yearid", sh.handleGetAttendanceDays)
	sh.teacherApp.Post("/attendancedays", sh.handleSyncAttendanceDays)
}

func (sh *SyncHandler) getTeacherAndYearID(ctx *fiber.Ctx) (teacherID, yearID string, ok bool) {
	teacherID, ok = ctx.Locals("id").(string)
	yearID = ctx.Params("yearid")
	if len(teacherID) == 0 || len(yearID) == 0 {
		ok = false
	}
	return
}

func (sh *SyncHandler) handleGetStudents(ctx *fiber.Ctx) error {
	teacherID, yearID, ok := sh.getTeacherAndYearID(ctx)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	students, err := sh.repo.GetStudents(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(students)
}

func (sh *SyncHandler) handleSyncStudents(ctx *fiber.Ctx) error {
	var students []common.AppStudentTx
	err := ctx.BodyParser(&students)
	if err != nil {
		return err
	}
	err = sh.repo.SyncStudents(students)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (sh *SyncHandler) handleGetClassPeriods(ctx *fiber.Ctx) error {
	teacherID, yearID, ok := sh.getTeacherAndYearID(ctx)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	classPeriods, err := sh.repo.GetClassPeriods(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(classPeriods)
}

func (sh *SyncHandler) handleSyncClassPeriods(ctx *fiber.Ctx) error {
	var classPeriods []common.AppClassPeriodTx
	err := ctx.BodyParser(&classPeriods)
	if err != nil {
		return err
	}
	err = sh.repo.SyncClassPeriods(classPeriods)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (sh *SyncHandler) handleGetAttendanceDays(ctx *fiber.Ctx) error {
	teacherID, yearID, ok := sh.getTeacherAndYearID(ctx)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	attendanceDays, err := sh.repo.GetAttendanceDays(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(attendanceDays)
}

func (sh *SyncHandler) handleSyncAttendanceDays(ctx *fiber.Ctx) error {
	var attendanceDays []common.AppAttendanceDayTx
	err := ctx.BodyParser(&attendanceDays)
	if err != nil {
		return err
	}
	err = sh.repo.SyncAttendanceDays(attendanceDays)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
