package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/teacher"
	"golang.org/x/crypto/bcrypt"
)

// SignUpHandler godoc
// @Summary registers a new teacher
// @Accept  json
// @Produce json
// @Param   teacher body     handlers.SignUpHandler.req true "teacher signup data"
// @Success 200     {object} handlers.SignUpHandler.res
// @Router  /teacher/signup [post]
func SignUpHandler(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type req struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type res struct {
		Message string `json:"message"`
	}
	return func(c *fiber.Ctx) error {
		var reqData req
		err := c.BodyParser(&reqData)
		if err != nil {
			return err
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		_, err = db.Teacher.Query().Where(teacher.EmailEQ(reqData.Email)).First(c.Context())
		// err nil means that teacher exists
		if err == nil {
			return NewClientErr(fiber.StatusBadRequest, "Ups! ya existe una cuenta con ese correo.")
		}
		_, err = db.Teacher.
			Create().
			SetID(newID()).
			SetName(reqData.Name).
			SetEmail(reqData.Email).
			SetLastName(reqData.LastName).
			SetPassword(string(hashedPassword)).
			Save(c.Context())
		if err != nil {
			return err
		}
		return c.SendStatus(http.StatusCreated)
	}
}
