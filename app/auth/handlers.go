package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
)

type AuthHandler struct {
	app        *fiber.App
	teacherApp fiber.Router
	adminApp   fiber.Router
	repo       AuthRepository
	config     config.Config
}

func NewAuthHandler(
	app *fiber.App,
	teacherApp fiber.Router,
	adminApp fiber.Router,
	repo AuthRepository,
	config config.Config,
) AuthHandler {
	return AuthHandler{app, teacherApp, adminApp, repo, config}
}

func (ah *AuthHandler) handle() {
	ah.app.Post("/signup", ah.HandleSignUp)
	ah.app.Post("/signin", ah.HandleSignIn)
	ah.teacherApp.Get("/profile", ah.HandleGetProfile)
	ah.adminApp.Get("/teachers/:id", ah.HandleGetTeacherProfile)
	ah.adminApp.Get("/teachers", ah.HandleGetTeachers)
}

type SignUpReq struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ah *AuthHandler) HandleSignUp(ctx *fiber.Ctx) error {
	var reqData SignUpReq
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	err = ah.repo.Register(reqData.Name, reqData.LastName, reqData.Email, reqData.Password)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignInRes struct {
	Teacher models.TeacherWithSubs `json:"teacher"`
	JWT     string                 `json:"jwt"`
}

func (ah *AuthHandler) HandleSignIn(ctx *fiber.Ctx) error {
	var reqData SignInReq
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	teacher, token, err := ah.repo.GetTeacher(reqData.Email, reqData.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(SignInRes{
		Teacher: teacher,
		JWT:     token,
	})
}

func (ah *AuthHandler) HandleGetProfile(ctx *fiber.Ctx) error {
	teacherID, ok := ctx.Locals("id").(string)
	if !ok {
		return errors.New("id not found in locals")
	}
	if len(teacherID) == 0 {
		return errors.New("id is empty")
	}
	profile, err := ah.repo.GetTeacherProfile(teacherID)
	if err != nil {
		return err
	}
	return ctx.JSON(profile)
}

func (ah *AuthHandler) HandleGetTeacherProfile(ctx *fiber.Ctx) error {
	teacherID := ctx.Params("id")
	profile, err := ah.repo.GetTeacherProfile(teacherID)
	if err != nil {
		return err
	}
	return ctx.JSON(profile)
}

func (ah *AuthHandler) HandleGetTeachers(ctx *fiber.Ctx) error {
	teachers, err := ah.repo.GetTeachers()
	if err != nil {
		return err
	}
	return ctx.JSON(teachers)
}
