package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/rigelapi/ent"
)

func AddSubscription(db *ent.Client, newID func() string) func(*fiber.Ctx) error {
	type req struct {
		TeacherID string `json:"teacher_id"`
		YearID    string `json:"year_id"`
		Method    string `json:"method"`
		Qtty      int    `json:"qtty"`
	}
	return func(c *fiber.Ctx) error {
		var reqData req
		err := c.BodyParser(&reqData)
		if err != nil {
			return err
		}
		_, err = db.Subscription.
			Create().
			SetID(newID()).
			SetMethod(reqData.Method).
			SetQtty(reqData.Qtty).
			SetDate(time.Now()).
			SetTeacherID(reqData.TeacherID).
			SetYearID(reqData.YearID).
			Save(c.Context())
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusCreated)
	}
}

func UpdateSubscription(db *ent.Client) func(*fiber.Ctx) error {
	type req struct {
		Method string `json:"method"`
		Qtty   int    `json:"qtty"`
	}
	return func(c *fiber.Ctx) error {
		subsID := c.Params("subscription_id")
		var reqData req
		err := c.BodyParser(&reqData)
		if err != nil {
			return err
		}
		_, err = db.Subscription.UpdateOneID(subsID).SetMethod(reqData.Method).SetQtty(reqData.Qtty).Save(c.Context())
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func DeleteSubscription(db *ent.Client) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		subsID := c.Params("subscription_id")
		err := db.Subscription.DeleteOneID(subsID).Exec(c.Context())
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
