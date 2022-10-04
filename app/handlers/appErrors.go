package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

type AppError struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Cause      string `json:"cause"`
	ErrorMsg   string `json:"error_msg"`
	ErrorStack string `json:"error_stack"`
}

func SaveAppErrors(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		appErrs := []AppError{}
		err := c.BodyParser(&appErrs)
		if err != nil {
			return err
		}
		toAdd := make([]*ent.AppErrorCreate, len(appErrs))
		for index, appErr := range appErrs {
			toAdd[index] = db.AppError.
				Create().
				SetID(appErr.ID).
				SetUserID(appErr.UserID).
				SetCause(appErr.Cause).
				SetErrorMsg(appErr.ErrorMsg).
				SetErrorStack(appErr.ErrorStack)
		}
		_, err = db.AppError.CreateBulk(toAdd...).Save(c.Context())
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
