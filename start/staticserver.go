package main

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
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


func CreateFile(name string) *os.File {
	file, _ := os.Create(name)
	return file
}

func RunServer() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/deps", func(c *fiber.Ctx) error {
		deps := []Dpto{}
		if err := c.BodyParser(&deps); err != nil {
			return err
		}
		file := CreateFile("departamentos.json")
		defer file.Close()

		depsJson, err := json.Marshal(deps)
		if err != nil {
			return err
		}
		toWrite := bytes.Buffer{}
		json.Indent(&toWrite, depsJson, "", "  ")
		file.Write(toWrite.Bytes())
		return c.JSON(deps)
	})

	app.Post("/provs", func(c *fiber.Ctx) error {
		provs := []Prov{}
		if err := c.BodyParser(&provs); err != nil {
			return err
		}
		file := CreateFile("provincias.json")
		defer file.Close()

		provsJson, err := json.Marshal(provs)
		if err != nil {
			return err
		}
		toWrite := bytes.Buffer{}
		json.Indent(&toWrite, provsJson, "", "  ")
		file.Write(toWrite.Bytes())
		return c.JSON(provs)
	})
	app.Post("/muns", func(c *fiber.Ctx) error {
		muns := []Mun{}
		if err := c.BodyParser(&muns); err != nil {
			return err
		}
		file := CreateFile("municipios.json")
		defer file.Close()

		munsJson, err := json.Marshal(muns)
		if err != nil {
			return err
		}
		toWrite := bytes.Buffer{}
		json.Indent(&toWrite, munsJson, "", "  ")
		file.Write(toWrite.Bytes())
		return c.JSON(muns)
	})
	app.Post("/cols", func(c *fiber.Ctx) error {
		cols := []School{}
		if err := c.BodyParser(&cols); err != nil {
			return err
		}
		file := CreateFile("colegios.json")
		defer file.Close()

		colsJson, err := json.Marshal(cols)
		if err != nil {
			return err
		}
		toWrite := bytes.Buffer{}
		json.Indent(&toWrite, colsJson, "", "  ")
		file.Write(toWrite.Bytes())
		return c.JSON(cols)
	})

	app.Listen(":4000")
}
