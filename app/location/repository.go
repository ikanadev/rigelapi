package location

import "github.com/vmkevv/rigelapi/app/models"

type LocationRepository interface {
	GetDeps() ([]models.Dpto, error)
}
