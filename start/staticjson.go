package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/vmkevv/rigelapi/ent"
)

type Dpto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Prov struct {
	ID    string `json:"id"`
	DepID string `json:"dpto_id"`
	Name  string `json:"name"`
}
type Mun struct {
	ID     string `json:"id"`
	ProvID string `json:"prov_id"`
	Name   string `json:"name"`
}
type School struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Lat   string `json:"lat"`
	Lon   string `json:"lon"`
	MunID string `json:"mun_id"`
}

func populateDepartamentos(client *ent.Client, ctx context.Context) error {
	content, err := ioutil.ReadFile("staticdata/departamentos.json")
	if err != nil {
		return err
	}
	fileDeps := []Dpto{}
	err = json.Unmarshal(content, &fileDeps)
	if err != nil {
		return err
	}
	dbDeps, err := client.Dpto.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.DptoCreate{}
	for _, fileDep := range fileDeps {
		exists := false
		for _, dbDep := range dbDeps {
			if fileDep.ID == dbDep.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Dpto.Create().SetID(fileDep.ID).SetName(fileDep.Name))
		}
	}
	_, err = client.Dpto.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateProvincias(client *ent.Client, ctx context.Context) error {
	content, err := ioutil.ReadFile("staticdata/provincias.json")
	if err != nil {
		return err
	}
	fileProvs := []Prov{}
	err = json.Unmarshal(content, &fileProvs)
	if err != nil {
		return err
	}
	dbProvs, err := client.Provincia.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.ProvinciaCreate{}
	for _, fileProv := range fileProvs {
		exists := false
		for _, dbProv := range dbProvs {
			if fileProv.ID == dbProv.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Provincia.Create().SetID(fileProv.ID).SetName(fileProv.Name).SetDepartamentoID(fileProv.DepID))
		}
	}
	_, err = client.Provincia.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateMunicipios(client *ent.Client, ctx context.Context) error {
	content, err := ioutil.ReadFile("staticdata/municipios.json")
	if err != nil {
		return err
	}
	fileMuns := []Mun{}
	err = json.Unmarshal(content, &fileMuns)
	if err != nil {
		return err
	}
	dbMuns, err := client.Municipio.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.MunicipioCreate{}
	for _, fileMun := range fileMuns {
		exists := false
		for _, dbMun := range dbMuns {
			if fileMun.ID == dbMun.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Municipio.Create().SetID(fileMun.ID).SetName(fileMun.Name).SetProvinciaID(fileMun.ProvID))
		}
	}
	_, err = client.Municipio.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateColegios(client *ent.Client, ctx context.Context) error {
	content, err := ioutil.ReadFile("staticdata/colegios.json")
	if err != nil {
		return err
	}
	fileCols := []School{}
	err = json.Unmarshal(content, &fileCols)
	if err != nil {
		return err
	}
	dbCols, err := client.School.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.SchoolCreate{}
	for _, fileCol := range fileCols {
		exists := false
		for _, dbCol := range dbCols {
			if fileCol.ID == dbCol.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(
				toSave,
				client.School.Create().SetID(fileCol.ID).SetName(fileCol.Name).SetLat(fileCol.Lat).SetLon(fileCol.Lon).SetMunicipioID(fileCol.MunID),
			)
		}
	}
	err = nil
	// postgres supports a max of 65535 parameters (5 params per insert)
	maxParameters := (1 << 16) - 1
	maxInserts := maxParameters / 5
	if len(toSave) > maxInserts {
		from := 0
		to := maxInserts
		for from != to {
			subArr := toSave[from:to]
			_, err = client.School.CreateBulk(subArr...).Save(ctx)
			if err != nil {
				break
			}
			from = to
			newTo := to + maxInserts
			if newTo > len(toSave) {
				to = len(toSave)
			} else {
				to = newTo
			}
		}
	} else {
		_, err = client.School.CreateBulk(toSave...).Save(ctx)
	}
  return err
}

func PopulateStaticJsonData(client *ent.Client, ctx context.Context) error {
	if err := populateDepartamentos(client, ctx); err != nil {
		return err
	}
	if err := populateProvincias(client, ctx); err != nil {
		return err
	}
	if err := populateMunicipios(client, ctx); err != nil {
		return err
	}
	if err := populateColegios(client, ctx); err != nil {
		return err
	}
	return nil
}
