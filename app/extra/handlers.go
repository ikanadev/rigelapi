package extra

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
)

type ExtraHandler struct {
	app  *fiber.App
	repo ExtraRepository
}

func NewExtraHandler(app *fiber.App, repo ExtraRepository) ExtraHandler {
	return ExtraHandler{app, repo}
}

func (eh *ExtraHandler) handle() {
	eh.app.Get("/years", eh.handleYearsData)
	eh.app.Get("/static", eh.handleStaticData)
	eh.app.Post("/errors", eh.handleSaveAppErrors)
	eh.app.Get("/stats", eh.handleStats)
	eh.app.Get("/class/:classid", eh.handleClassDetails)
}

func (eh *ExtraHandler) handleYearsData(ctx *fiber.Ctx) error {
	years, err := eh.repo.GetYearsData()
	if err != nil {
		return err
	}
	return ctx.JSON(years)
}

func (eh *ExtraHandler) handleStaticData(ctx *fiber.Ctx) error {
	grades, err := eh.repo.GetGrades()
	if err != nil {
		return err
	}
	subjects, err := eh.repo.GetSubjects()
	if err != nil {
		return err
	}
	return ctx.JSON(StaticDataRes{grades, subjects})
}

func (eh *ExtraHandler) handleSaveAppErrors(ctx *fiber.Ctx) error {
	appErrs := []models.AppError{}
	err := ctx.BodyParser(&appErrs)
	if err != nil {
		return err
	}
	err = eh.repo.SaveAppErrors(appErrs)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (eh *ExtraHandler) handleStats(ctx *fiber.Ctx) error {
	teachers, err := eh.repo.GetTeachersCount()
	if err != nil {
		return err
	}
	classes, err := eh.repo.GetClassesCount()
	if err != nil {
		return err
	}
	schools, err := eh.repo.GetClassesCount()
	if err != nil {
		return err
	}
	acts, err := eh.repo.GetActivitiesCount()
	if err != nil {
		return err
	}
	return ctx.JSON(StatsRes{
		Stats: Stats{
			Teachers:   teachers,
			Classes:    classes,
			Schools:    schools,
			Activities: acts,
		},
	})
}

func (eh *ExtraHandler) handleClassDetails(ctx *fiber.Ctx) error {
	classID := ctx.Params("classid")
	var resp ClassDetailsResp
	classData, err := eh.repo.GetClassData(classID)
	if err != nil {
		return err
	}
	resp.ClassData = classData

	classPeriods, err := eh.repo.GetClassPeriodsData(classID)
	if err != nil {
		return err
	}
	resp.ClassPeriods = classPeriods

	students, err := eh.repo.GetStudentsData(classID, classPeriods)
	if err != nil {
		return err
	}
	resp.Students = students
	return ctx.JSON(resp)
}

type StaticDataRes struct {
	Grades   []models.Grade   `json:"grades"`
	Subjects []models.Subject `json:"subjects"`
}

type Stats struct {
	Teachers   int `json:"teachers"`
	Classes    int `json:"classes"`
	Schools    int `json:"schools"`
	Activities int `json:"activities"`
}
type StatsRes struct {
	Stats Stats `json:"stats"`
}
type ClassDetailsResp struct {
	ClassData    models.ClassData         `json:"class_data"`
	Students     []models.StudentData     `json:"students"`
	ClassPeriods []models.ClassPeriodData `json:"class_periods"`
}
