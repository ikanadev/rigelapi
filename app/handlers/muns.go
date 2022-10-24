package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

func MunsHandler(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		provID := c.Params("provid")
		muns, err := db.Municipio.Query().Where(municipio.HasProvinciaWith(provincia.ID(provID))).Order(ent.Asc(municipio.FieldName)).All(c.Context())
		if err != nil {
			return err
		}
		munsResp := make([]Municipio, len(muns))
		for i, mun := range muns {
			munsResp[i] = Municipio{mun.ID, mun.Name}
		}
		return c.JSON(munsResp)
	}
}
