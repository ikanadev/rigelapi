package app

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmkevv/rigelapi/app/handlers"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
	"github.com/vmkevv/rigelapi/utils"
)

type Server struct {
	DB           *ent.Client
	Config       config.Config
	App          *fiber.App
	ProtectedApp fiber.Router
	AdminApp     fiber.Router
	IDGenerator  func() string
	DBCtx        context.Context
}
type errMsg struct {
	Message string `json:"message"`
}

func NewServer(db *ent.Client, config config.Config, logger *log.Logger, dbCtx context.Context) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "Unespected internal server error"
			if err, ok := err.(handlers.ClientErr); ok {
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
	return Server{
		db,
		config,
		app,
		app.Group("/auth", authMiddleware(config)),
		app.Group("/admin", authMiddleware(config), adminMiddleware(db)),
		utils.NanoIDGenerator(),
		dbCtx,
	}
}

func (server Server) Run() {
	server.App.Use(cors.New())
	server.ProtectedApp.Post("/class", handlers.NewClassHandler(server.DB, server.IDGenerator))
	server.ProtectedApp.Post("/students", handlers.SaveStudent(server.DB))
	server.ProtectedApp.Get("/students/year/:yearid", handlers.GetStudents(server.DB))
	server.ProtectedApp.Post("/classperiods", handlers.SaveClassPeriods(server.DB))
	server.ProtectedApp.Get("/classperiods/year/:yearid", handlers.GetClassPeriods(server.DB))
	server.ProtectedApp.Post("/attendancedays", handlers.SaveAttendanceDays(server.DB))
	server.ProtectedApp.Get("/attendancedays/year/:yearid", handlers.GetAttendanceDays(server.DB))
	server.ProtectedApp.Post("/attendances", handlers.SaveAttendances(server.DB))
	server.ProtectedApp.Get("/attendances/year/:yearid", handlers.GetAttendances(server.DB))
	server.ProtectedApp.Post("/activities", handlers.SaveActivities(server.DB))
	server.ProtectedApp.Get("/activities/year/:yearid", handlers.GetActivities(server.DB))
	server.ProtectedApp.Post("/scores", handlers.SaveScores(server.DB))
	server.ProtectedApp.Get("/scores/year/:yearid", handlers.GetScores(server.DB))
	server.AdminApp.Get("/teachers", handlers.GetTeachers(server.DB))
	server.AdminApp.Get("/teacher/:id", handlers.GetTeacher(server.DB))
	server.AdminApp.Post("/subscription", handlers.AddSubscription(server.DB, server.IDGenerator))
	server.AdminApp.Patch("/subscription/:subscription_id", handlers.UpdateSubscription(server.DB))
	server.AdminApp.Delete("/subscription/:subscription_id", handlers.DeleteSubscription(server.DB))
}
