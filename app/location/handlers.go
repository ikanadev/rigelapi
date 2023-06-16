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
	lh.app.Get("/provs/dep/:depid", lh.handleGetProvs)
}

func (lh *LocationHandler) handleGetDeps(ctx *fiber.Ctx) error {
	deps, err := lh.repo.GetDeps()
	if err != nil {
		return err
	}
	return ctx.JSON(deps)
}

func (lh *LocationHandler) handleGetProvs(ctx *fiber.Ctx) error {
	depId := ctx.Params("depid")
	provs, err := lh.repo.GetProvs(depId)
	if err != nil {
		return err
	}
	return ctx.JSON(provs)
}
