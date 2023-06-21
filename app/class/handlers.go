package class

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
)

type ClassHandler struct {
	app  *fiber.App
	repo ClassRepository
}

func NewClassHandler(app *fiber.App, repo ClassRepository) ClassHandler {
	return ClassHandler{app, repo}
}

func (ch *ClassHandler) handle() {
	ch.app.Get("/class/:classid", ch.HandleClassDetails)
}

func (ch *ClassHandler) HandleClassDetails(ctx *fiber.Ctx) error {
	classID := ctx.Params("classid")
	var resp ClassDetailsResp
	classData, err := ch.repo.GetClassData(classID)
	if err != nil {
		return err
	}
	resp.ClassData = classData

	classPeriods, err := ch.repo.GetClassPeriodsData(classID)
	if err != nil {
		return err
	}
	resp.ClassPeriods = classPeriods

	students, err := ch.repo.GetStudentsData(classID, classPeriods)
	if err != nil {
		return err
	}
	resp.Students = students
	return ctx.JSON(resp)
}

type ClassDetailsResp struct {
	ClassData    models.ClassData         `json:"class_data"`
	Students     []models.StudentData     `json:"students"`
	ClassPeriods []models.ClassPeriodData `json:"class_periods"`
}
