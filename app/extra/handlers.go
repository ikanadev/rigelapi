package extra

import (
	"github.com/gofiber/fiber/v2"
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
}

func (eh *ExtraHandler) handleYearsData(ctx *fiber.Ctx) error {
	years, err := eh.repo.GetYearsData()
	if err != nil {
		return err
	}
	return ctx.JSON(years)
}
