package teacher

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type TeacherHandler struct {
	teacherApp fiber.Router
	repo       TeacherRepository
}

func NewTeacherHandler(
	teacherApp fiber.Router,
	repo TeacherRepository,
) TeacherHandler {
	return TeacherHandler{teacherApp, repo}
}

func (ah *TeacherHandler) handle() {
	ah.teacherApp.Get("/profile", ah.HandleGetProfile)
}

func (ah *TeacherHandler) HandleGetProfile(ctx *fiber.Ctx) error {
	teacherID, ok := ctx.Locals("id").(string)
	if !ok {
		return errors.New("id not found in locals")
	}
	if len(teacherID) == 0 {
		return errors.New("id is empty")
	}
	profile, err := ah.repo.GetProfile(teacherID)
	if err != nil {
		return err
	}
	return ctx.JSON(profile)
}
