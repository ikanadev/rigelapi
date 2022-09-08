package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/vmkevv/rigelapi/app/handlers"
	"github.com/vmkevv/rigelapi/config"
	_ "github.com/vmkevv/rigelapi/docs"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/utils"
)

type Server struct {
	db     *ent.Client
	config config.Config
	app    *fiber.App
	newID  func() string
}
type errMsg struct {
	Message string `json:"message"`
}

func NewServer(db *ent.Client, config config.Config) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "Unespected internal server error"
			if err, ok := err.(handlers.ClientErr); ok {
				code = err.Status
				msg = err.Error()
			}
			if code == fiber.StatusInternalServerError {
				log.Printf("Fatal: %v\n", err)
			}
			return c.Status(code).JSON(errMsg{Message: msg})
		},
	})
	return Server{
		db,
		config,
		app,
		utils.NanoIDGenerator(),
	}
}

func (server Server) Run() {
	server.app.Use(cors.New())
	// Swagger handler
	server.app.Get("/swagger/*", swagger.HandlerDefault)
	server.app.Post("/signup", handlers.SignUpHandler(server.db, server.newID))
	server.app.Post("/signin", handlers.SignInHandler(server.db, server.config))

	protected := server.app.Group("/api", authMiddleware(server.config))
	protected.Get("/deps", handlers.DepsHandler(server.db))
	protected.Get("/provs/dep/:depid", handlers.ProvsHandler(server.db))

	server.app.Listen(fmt.Sprintf(":%s", server.config.App.Port))
}
