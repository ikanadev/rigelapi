package subscription

import "github.com/gofiber/fiber/v2"

type SubscriptionHandler struct {
	adminApp fiber.Router
	repo     SubscriptionRepository
}

func NewSubscriptionHandler(
	adminApp fiber.Router,
	repo SubscriptionRepository,
) SubscriptionHandler {
	return SubscriptionHandler{adminApp, repo}
}

func (sh SubscriptionHandler) handle() {
	sh.adminApp.Post("/subscriptions", sh.handleSaveSubscription)
	sh.adminApp.Patch("/subscriptions/:id", sh.handleUpdateSubscription)
	sh.adminApp.Delete("/subscriptions/:id", sh.handleDeleteSubscription)
}

type SaveSubscriptionReq struct {
	TeacherID string `json:"teacher_id"`
	YearID    string `json:"year_id"`
	Method    string `json:"method"`
	Qtty      int    `json:"qtty"`
}

func (sh SubscriptionHandler) handleSaveSubscription(ctx *fiber.Ctx) error {
	var reqData SaveSubscriptionReq
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	err = sh.repo.SaveSubscription(reqData.TeacherID, reqData.YearID, reqData.Method, reqData.Qtty)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusCreated)
}

type UpdateSubscriptionReq struct {
	Method string `json:"method"`
	Qtty   int    `json:"qtty"`
}

func (sh SubscriptionHandler) handleUpdateSubscription(ctx *fiber.Ctx) error {
	var reqData UpdateSubscriptionReq
	subID := ctx.Params("id")
	err := ctx.BodyParser(&reqData)
	if err != nil {
		return err
	}
	err = sh.repo.UpdateSubscription(subID, reqData.Method, reqData.Qtty)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (sh SubscriptionHandler) handleDeleteSubscription(ctx *fiber.Ctx) error {
	subID := ctx.Params("id")
	err := sh.repo.DeleteSubscription(subID)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}
