package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
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
      return c.Status(500).JSON(newErrMsg("Bad request data"))
    }
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.MinCost)
    if err != nil {
      return c.Status(http.StatusInternalServerError).JSON(newErrMsg("Error hashing password"))
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
      return c.Status(http.StatusInternalServerError).JSON(newErrMsg("Error saving teacher data"))
    }
		return c.SendStatus(http.StatusCreated)
	}
}
