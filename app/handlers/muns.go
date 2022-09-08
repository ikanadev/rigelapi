package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// MunsHandler godoc
// @Summary List all municipios from prov with {provid}
// @Produce json
// @Param   provid path     int true "Provincia ID"
// @Success 200    {object} handlers.MunsHandler.res
// @Router  /muns/prov/{provid} [get]
func MunsHandler(db *ent.Client) func(*fiber.Ctx) error {
	type res = []*ent.Municipio
	return func(c *fiber.Ctx) error {
		provID := c.Params("provid")
		muns, err := db.Municipio.Query().Where(municipio.HasProvinciaWith(provincia.ID(provID))).All(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(muns)
	}
}
