package app

import (
	"fmt"

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
	db      *ent.Client
	config  config.Config
	app     *fiber.App
	newID func() string
}

func NewServer(db *ent.Client, config config.Config) Server {
  type util struct {}
	return Server{
    db,
    config,
    fiber.New(),
    utils.NanoIDGenerator(),
  }
}

func (server Server) Run() {
	server.app.Use(cors.New())
	// Swagger handler
	server.app.Get("/swagger/*", swagger.HandlerDefault)
	server.app.Post("/teacher/signup", handlers.SignUpHandler(server.db, server.newID))
	// TODO: serve routes here
	server.app.Listen(fmt.Sprintf(":%s", server.config.App.Port))
}
