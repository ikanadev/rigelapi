package extra

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
)

type ExtraEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewExtraEntRepo(ent *ent.Client, ctx context.Context) ExtraEntRepo {
	return ExtraEntRepo{ent, ctx}
}

func (eer ExtraEntRepo) GetYearsData() ([]models.YearData, error) {
	entYears, err := eer.ent.Year.Query().WithAreas().WithPeriods().All(eer.ctx)
	if err != nil {
		return nil, err
	}
	years := make([]models.YearData, len(entYears))
	for i, year := range entYears {
		periods := make([]models.Period, len(year.Edges.Periods))
		for j, period := range year.Edges.Periods {
			periods[j] = models.Period{ID: period.ID, Name: period.Name}
		}
		areas := make([]models.Area, len(year.Edges.Areas))
		for j, area := range year.Edges.Areas {
			areas[j] = models.Area{
				ID:     area.ID,
				Name:   area.Name,
				Points: area.Points,
			}
		}

		years[i] = models.YearData{
			Year: models.Year{
				ID:    year.ID,
				Value: year.Value,
			},
			Periods: periods,
			Areas:   areas,
		}
	}

	return years, nil
}
