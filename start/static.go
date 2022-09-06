package main

import (
	"context"

	"github.com/vmkevv/rigelapi/ent"
)

func populateSubjects(client *ent.Client, ctx context.Context) error {
	subjectsMap := map[string]string{
		"1":  "BIOLOGÍA - GEOGRAFÍA",
		"2":  "FÍSICA - QUÍMICA",
		"3":  "LENGUA CASTELLANA Y ORIGINARIA",
		"4":  "LENGUA EXTRANJERA",
		"5":  "CIENCIAS SOCIALES",
		"6":  "EDUCACIÓN FÍSICA Y DEPORTES",
		"7":  "EDUCACIÓN MUSICAL",
		"8":  "ARTES PLÁSTICAS Y VISUALES",
		"9":  "COSMOVISIONES, FILOSOFÍA Y PSICOLOGÍA",
		"10": "VALORES, ESPIRITUALIDAD Y RELIGIONES",
		"11": "MATEMÁTICA",
		"12": "TÉCNICA Y TECNOLOGÍA",
	}
	dbSubjects, err := client.Subject.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.SubjectCreate{}
	for id, subject := range subjectsMap {
		exists := false
		for _, dbSubject := range dbSubjects {
			if id == dbSubject.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Subject.Create().SetID(id).SetName(subject))
		}
	}
	_, err = client.Subject.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateGrades(client *ent.Client, ctx context.Context) error {
	gradesMap := map[string]string{
		"1": "1RO DE SECUNDARIA",
		"2": "2DO DE SECUNDARIA",
		"3": "3RO DE SECUNDARIA",
		"4": "4TO DE SECUNDARIA",
		"5": "5TO DE SECUNDARIA",
		"6": "6TO DE SECUNDARIA",
	}
	dbGrades, err := client.Grade.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.GradeCreate{}
	for id, grade := range gradesMap {
		exists := false
		for _, dbGrade := range dbGrades {
			if id == dbGrade.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Grade.Create().SetID(id).SetName(grade))
		}
	}
	_, err = client.Grade.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateYears(client *ent.Client, ctx context.Context) error {
	yearsMap := map[string]int{
		"1": 2022,
		"2": 2023,
	}
	dbYears, err := client.Year.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.YearCreate{}
	for id, year := range yearsMap {
		exists := false
		for _, dbYear := range dbYears {
			if id == dbYear.ID {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Year.Create().SetID(id).SetValue(year))
		}
	}
	_, err = client.Year.CreateBulk(toSave...).Save(ctx)
  return err
}

func populateAreas(client *ent.Client, ctx context.Context) error {
	type Areas struct {
		id     string
		name   string
		points int
		yearId string
	}
	staticAreas := []Areas{
		{"1", "Ser", 10, "1"},
		{"2", "Saber", 35, "1"},
		{"3", "Hacer", 35, "1"},
		{"4", "Decidir", 10, "1"},
		{"5", "Autoevaluación", 10, "1"},
	}
	dbAreas, err := client.Area.Query().All(ctx)
	if err != nil {
		return err
	}
	toSave := []*ent.AreaCreate{}
	for _, staticArea := range staticAreas {
		exists := false
		for _, dbArea := range dbAreas {
			if dbArea.ID == staticArea.id {
				exists = true
				break
			}
		}
		if !exists {
			toSave = append(toSave, client.Area.Create().SetID(staticArea.id).SetName(staticArea.name).SetPoints(staticArea.points).SetYearID(staticArea.yearId))
		}
	}
	_, err = client.Area.CreateBulk(toSave...).Save(ctx)
  return err
}

func populatePeriods(client *ent.Client, ctx context.Context) error {
  return nil
}

func PopulateStaticData(client *ent.Client, ctx context.Context) error {
	if err := populateSubjects(client, ctx); err != nil {
		return err
	}
	if err := populateGrades(client, ctx); err != nil {
		return err
	}
	if err := populateYears(client, ctx); err != nil {
		return err
	}
	if err := populateAreas(client, ctx); err != nil {
		return err
	}
	return nil
}
