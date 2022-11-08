package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

type Resp map[string][][]string

func ParseXLS() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		formFile, err := c.FormFile("xls")
		if err != nil {
			return err
		}
		file, err := formFile.Open()
		if err != nil {
			return err
		}
		xlsFile, err := excelize.OpenReader(file)
		if err != nil {
			return err
		}
		sheets := xlsFile.GetSheetList()
		resp := Resp{}
		for _, sheetName := range sheets {
			rows, err := xlsFile.GetRows(sheetName)
			if err != nil {
				continue
			}
			resp[sheetName] = rows
		}

		return c.JSON(resp)
	}
}
