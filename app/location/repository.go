package location

import "github.com/vmkevv/rigelapi/app/models"

type LocationRepository interface {
	GetDeps() ([]models.Dpto, error)
	GetProvs(depID string) ([]models.Provincia, error)
	GetMuns(provID string) ([]models.Municipio, error)
}
