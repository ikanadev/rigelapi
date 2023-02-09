package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func readDeps() []string {
	content, err := ioutil.ReadFile("departamentos.json")
	if err != nil {
		panic(err)
	}
	deps := []Dpto{}
	err = json.Unmarshal(content, &deps)
	if err != nil {
		panic(err)
	}
	var sql []string = []string{"INSERT INTO dptos (id, name) VALUES\n"}
	for i, dep := range deps {
		separator := ","
		if i == len(deps)-1 {
			separator = ";"
		}
		sql = append(sql, fmt.Sprintf("('%s', '%s')%s\n", dep.ID, dep.Name, separator))
	}
	return sql
}
func readProvs() []string {
	content, err := ioutil.ReadFile("provincias.json")
	if err != nil {
		panic(err)
	}
	provs := []Prov{}
	err = json.Unmarshal(content, &provs)
	if err != nil {
		panic(err)
	}
	var sql []string = []string{"INSERT INTO provincia (id, name, dpto_provincias) VALUES\n"}
	for i, prov := range provs {
		separator := ","
		if i == len(provs)-1 {
			separator = ";"
		}
		sql = append(sql, fmt.Sprintf("('%s', '%s', '%s')%s\n", prov.ID, prov.Name, prov.DepID, separator))
	}
	return sql
}
func readMuns() []string {
	content, err := ioutil.ReadFile("municipios.json")
	if err != nil {
		panic(err)
	}
	muns := []Mun{}
	err = json.Unmarshal(content, &muns)
	if err != nil {
		panic(err)
	}
	var sql []string = []string{"INSERT INTO municipios (id, name, provincia_municipios) VALUES\n"}
	for i, mun := range muns {
		separator := ","
		if i == len(muns)-1 {
			separator = ";"
		}
		sql = append(sql, fmt.Sprintf("('%s', '%s', '%s')%s\n", mun.ID, mun.Name, mun.ProvID, separator))
	}
	return sql
}
func readSchools() []string {
	content, err := ioutil.ReadFile("colegios.json")
	if err != nil {
		panic(err)
	}
	schools := []School{}
	err = json.Unmarshal(content, &schools)
	if err != nil {
		panic(err)
	}
	var sql []string = []string{"INSERT INTO schools (id, name, lat, lon, municipio_schools) VALUES\n"}
	for i, school := range schools {
		separator := ","
		if i == len(schools)-1 {
			separator = ";"
		}
		sql = append(sql, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')%s\n", school.ID, school.Name, school.Lat, school.Lon, school.MunID, separator))
	}
	return sql
}

func main() {
	fileName := "sql.sql"
	deps := readDeps()
	provs := readProvs()
	muns := readMuns()
	schools := readSchools()

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	writter := bufio.NewWriter(file)

	for _, dep := range deps {
		writter.WriteString(dep)
	}
	writter.WriteString("\n")
	for _, prov := range provs {
		writter.WriteString(prov)
	}
	writter.WriteString("\n")
	for _, mun := range muns {
		writter.WriteString(mun)
	}
	writter.WriteString("\n")
	for _, school := range schools {
		writter.WriteString(school)
	}

	writter.Flush()

}
