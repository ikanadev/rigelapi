package location

import "github.com/gofiber/fiber/v2"

type LocationHandler struct {
	app  fiber.Router
	repo LocationRepository
}

func NewLocationHandler(app fiber.Router, repo LocationRepository) LocationHandler {
	return LocationHandler{app, repo}
}

func (lh *LocationHandler) handle() {
	lh.app.Get("/deps", lh.handleGetDeps)
	lh.app.Get("/provs/dep/:depid", lh.handleGetProvs)
	lh.app.Get("/muns/prov/:provid", lh.handleGetMuns)
	lh.app.Get("/schools/mun/:munid", lh.handleGetSchools)
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

func (lh *LocationHandler) handleGetMuns(ctx *fiber.Ctx) error {
	provID := ctx.Params("provid")
	muns, err := lh.repo.GetMuns(provID)
	if err != nil {
		return err
	}
	return ctx.JSON(muns)
}

func (lh *LocationHandler) handleGetSchools(ctx *fiber.Ctx) error {
	munID := ctx.Params("munid")
	schools, err := lh.repo.GetSchools(munID)
	if err != nil {
		return err
	}
	return ctx.JSON(schools)
}
