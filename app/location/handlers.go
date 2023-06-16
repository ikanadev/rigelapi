package location

import "github.com/gofiber/fiber/v2"

type LocationHandler struct {
	app  *fiber.App
	repo LocationRepository
}

func NewLocationHandler(app *fiber.App, repo LocationRepository) LocationHandler {
	return LocationHandler{app, repo}
}

func (lh *LocationHandler) handle() {
	lh.app.Get("/deps", lh.handleGetDeps)
}

func (lh *LocationHandler) handleGetDeps(ctx *fiber.Ctx) error {
	deps, err := lh.repo.GetDeps()
	if err != nil {
		return err
	}
	return ctx.JSON(deps)
}
