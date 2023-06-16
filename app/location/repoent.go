package location

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/dpto"
)

type LocationEntRepo struct {
	ent *ent.Client
	ctx context.Context
}

func NewLocationEntRepo(ent *ent.Client, ctx context.Context) LocationEntRepo {
	return LocationEntRepo{
		ent,
		ctx,
	}
}

func (ler LocationEntRepo) GetDeps() ([]models.Dpto, error) {
	entDeps, err := ler.ent.Dpto.Query().Order(ent.Asc(dpto.FieldName)).All(ler.ctx)
	if err != nil {
		return nil, err
	}
	deps := make([]models.Dpto, len(entDeps))
	for i, dep := range entDeps {
		deps[i] = models.Dpto{ID: dep.ID, Name: dep.Name}
	}
	return deps, nil
}
