package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func YearlyDataHandler(db *ent.Client) func(*fiber.Ctx) error {
	type YearResp struct {
		Year
		Periods []Period `json:"periods"`
		Areas   []Area   `json:"areas"`
	}

	return func(c *fiber.Ctx) error {
		years, err := db.Year.Query().WithAreas().WithPeriods().All(c.Context())
		if err != nil {
			return err
		}
		yearsResp := make([]YearResp, len(years))
		for i, year := range years {
			periodsResp := make([]Period, len(year.Edges.Periods))
			for j, period := range year.Edges.Periods {
				periodsResp[j] = Period{period.ID, period.Name}
			}
			areasResp := make([]Area, len(year.Edges.Areas))
			for j, area := range year.Edges.Areas {
				areasResp[j] = Area{area.ID, area.Name, area.Points}
			}

			yearsResp[i] = YearResp{
				Year: Year{
					ID:    year.ID,
					Value: year.Value,
				},
				Periods: periodsResp,
				Areas:   areasResp,
			}
		}
		return c.JSON(yearsResp)
	}
}
