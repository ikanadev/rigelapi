package app

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmkevv/rigelapi/app/common"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/utils"
)

type Server struct {
	DB          *ent.Client
	Config      config.Config
	App         *fiber.App
	TeacherApp  fiber.Router
	AdminApp    fiber.Router
	IDGenerator func() string
	DBCtx       context.Context
}
type errMsg struct {
	Message string `json:"message"`
}

func NewServer(db *ent.Client, config config.Config, logger *log.Logger, dbCtx context.Context) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "Unespected internal server error"
			if err, ok := err.(common.ClientErr); ok {
				code = err.Status
				msg = err.Error()
			}
			if code == fiber.StatusInternalServerError {
				log.Println(err.Error())
				userId, _ := c.Locals("id").(string)
				// Method;IP;BaseURL;Path;Protocol\n",
				logger.Printf(
					"%s|%s|%s|%s|%s|%s|%s\n",
					c.Method(),
					c.IP(),
					c.BaseURL(),
					c.Path(),
					userId,
					c.Get("User-Agent"),
					err.Error(),
				)
			}
			return c.Status(code).JSON(errMsg{Message: msg})
		},
	})
	app.Use(cors.New())
	return Server{
		DB:          db,
		Config:      config,
		App:         app,
		TeacherApp:  app.Group("/auth", authMiddleware(config)),
		AdminApp:    app.Group("/admin", authMiddleware(config), adminMiddleware(db)),
		IDGenerator: utils.NanoIDGenerator(),
		DBCtx:       dbCtx,
	}
}
