package class

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
)

type ClassHandler struct {
	app        *fiber.App
	teacherApp fiber.Router
	repo       ClassRepository
}

func NewClassHandler(
	app *fiber.App,
	teacherApp fiber.Router,
	repo ClassRepository,
) ClassHandler {
	return ClassHandler{app, teacherApp, repo}
}

func (ch *ClassHandler) handle() {
	ch.app.Get("/class/:classid", ch.HandleClassDetails)
	ch.teacherApp.Get("/classes/year/:yearid", ch.HandleTeacherClasses)
	ch.teacherApp.Post("/class", ch.HandleSaveClass)
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

func (ch *ClassHandler) HandleTeacherClasses(ctx *fiber.Ctx) error {
	teacherID := ctx.Locals("id").(string)
	yearID := ctx.Params("yearid")
	if len(yearID) == 0 {
		return errors.New("No yearid found")
	}
	classes, err := ch.repo.GetTeacherClasses(teacherID, yearID)
	if err != nil {
		return err
	}
	return ctx.JSON(classes)
}

func (ch *ClassHandler) HandleSaveClass(ctx *fiber.Ctx) error {
	teacherID, ok := ctx.Locals("id").(string)
	var reqData SaveClassReq
	if !ok || len(teacherID) == 0 {
		return errors.New("Cannot read string ID from context or it is empty")
	}
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	err = ch.repo.SaveClass(NewClassData{
		YearID:    reqData.YearID,
		TeacherID: teacherID,
		GradeID:   reqData.GradeID,
		SubjectID: reqData.SubjectID,
		SchoolID:  reqData.SchoolID,
		Parallel:  reqData.Parallel,
	})
	if err != nil {
		return err
	}
	classes, err := ch.repo.GetTeacherClasses(teacherID, reqData.YearID)
	if err != nil {
		return err
	}
	return ctx.JSON(classes)
}

type SaveClassReq struct {
	GradeID   string `json:"gradeId"`
	SubjectID string `json:"subjectId"`
	SchoolID  string `json:"schoolId"`
	YearID    string `json:"yearId"`
	Parallel  string `json:"parallel"`
}

type ClassDetailsResp struct {
	ClassData    models.ClassData         `json:"class_data"`
	Students     []models.StudentData     `json:"students"`
	ClassPeriods []models.ClassPeriodData `json:"class_periods"`
}
