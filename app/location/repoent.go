package location

import (
	"context"

	"github.com/vmkevv/rigelapi/app/models"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/provincia"
	"github.com/vmkevv/rigelapi/ent/school"
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

func (ler LocationEntRepo) GetProvs(depID string) ([]models.Provincia, error) {
	entProvs, err := ler.ent.Provincia.Query().
		Where(provincia.HasDepartamentoWith(dpto.ID(depID))).
		Order(ent.Asc(provincia.FieldName)).
		All(ler.ctx)
	if err != nil {
		return nil, err
	}
	provs := make([]models.Provincia, len(entProvs))
	for i, prov := range entProvs {
		provs[i] = models.Provincia{ID: prov.ID, Name: prov.Name}
	}
	return provs, nil
}

func (ler LocationEntRepo) GetMuns(provID string) ([]models.Municipio, error) {
	entMuns, err := ler.ent.Municipio.Query().
		Where(municipio.HasProvinciaWith(provincia.ID(provID))).
		Order(ent.Asc(municipio.FieldName)).
		All(ler.ctx)
	if err != nil {
		return nil, err
	}
	muns := make([]models.Municipio, len(entMuns))
	for i, mun := range entMuns {
		muns[i] = models.Municipio{ID: mun.ID, Name: mun.Name}
	}
	return muns, nil
}

func (ler LocationEntRepo) GetSchools(munID string) ([]models.School, error) {
	entSchools, err := ler.ent.School.Query().
		Where(school.HasMunicipioWith(municipio.ID(munID))).
		Order(ent.Asc(school.FieldName)).
		All(ler.ctx)
	if err != nil {
		return nil, err
	}
	schools := make([]models.School, len(entSchools))
	for i, school := range entSchools {
		schools[i] = models.School{
			ID:   school.ID,
			Name: school.Name,
			Lat:  school.Lat,
			Lon:  school.Lon,
		}
	}
	return schools, nil
}
