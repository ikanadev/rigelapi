package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/config"
)

type AuthHandler struct {
	app      fiber.Router
	authRepo AuthRepository
	config   config.Config
}

func NewAuthHandler(app fiber.Router, repo AuthRepository, config config.Config) AuthHandler {
	return AuthHandler{app, repo, config}
}

func (ah *AuthHandler) handle() {
	ah.app.Post("/signup", ah.HandleSignUp)
	ah.app.Post("/signin", ah.HandleSignIn)
}

func (ah *AuthHandler) HandleSignUp(ctx *fiber.Ctx) error {
	var reqData SignUpReq
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	err = ah.authRepo.Register(reqData.Name, reqData.LastName, reqData.Email, reqData.Password)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

func (ah *AuthHandler) HandleSignIn(ctx *fiber.Ctx) error {
	var reqData SignInReq
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	teacher, token, err := ah.authRepo.GetTeacher(reqData.Email, reqData.Password)
	if err != nil {
		return err
	}
	return ctx.JSON(SignInRes{
		Teacher: teacher,
		JWT:     token,
	})
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SignInRes struct {
	Teacher models.TeacherWithSubs `json:"teacher"`
	JWT     string                 `json:"jwt"`
}

type SignUpReq struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
