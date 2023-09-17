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
	sh.teacherApp.Get("/attendances/year/:yearid", sh.handleGetAttendances)
	sh.teacherApp.Post("/attendances", sh.handleSyncAttendances)
	sh.teacherApp.Get("/activities/year/:yearid", sh.handleGetActivities)
	sh.teacherApp.Post("/activities", sh.handleSyncActivities)
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

func (sh *SyncHandler) handleGetAttendances(ctx *fiber.Ctx) error {
	teacherID, yearID, ok := sh.getTeacherAndYearID(ctx)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	attendances, err := sh.repo.GetAttendances(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(attendances)
}

func (sh *SyncHandler) handleSyncAttendances(ctx *fiber.Ctx) error {
	var attendances []common.AppAttendanceTx
	err := ctx.BodyParser(&attendances)
	if err != nil {
		return err
	}
	err = sh.repo.SyncAttendances(attendances)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (sh *SyncHandler) handleGetActivities(ctx *fiber.Ctx) error {
	teacherID, yearID, ok := sh.getTeacherAndYearID(ctx)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).SendString("No autorizado")
	}
	activities, err := sh.repo.GetActivities(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(activities)
}

func (sh *SyncHandler) handleSyncActivities(ctx *fiber.Ctx) error {
	var activities []common.AppActivityTx
	err := ctx.BodyParser(&activities)
	if err != nil {
		return err
	}
	err = sh.repo.SyncActivities(activities)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
