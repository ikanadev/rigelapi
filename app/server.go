package app

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmkevv/rigelapi/app/handlers"
	"github.com/vmkevv/rigelapi/config"
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

func NewServer(db *ent.Client, config config.Config, logger *log.Logger) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "Unespected internal server error"
			if err, ok := err.(handlers.ClientErr); ok {
				code = err.Status
				msg = err.Error()
			}
			if code == fiber.StatusInternalServerError {
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
		utils.NanoIDGenerator(),
	}
}

func (server Server) Run() error {
	server.app.Use(cors.New())
	server.app.Use(logUserAgent())

	server.app.Post("/signup", handlers.SignUpHandler(server.db, server.newID))
	server.app.Post("/signin", handlers.SignInHandler(server.db, server.config))
	server.app.Get("/deps", handlers.DepsHandler(server.db))
	server.app.Get("/provs/dep/:depid", handlers.ProvsHandler(server.db))
	server.app.Get("/muns/prov/:provid", handlers.MunsHandler(server.db))
	server.app.Get("/schools/mun/:munid", handlers.SchoolsHandler(server.db))
	server.app.Get("/years", handlers.YearlyDataHandler(server.db))
	server.app.Get("/static", handlers.StaticDataHandler(server.db))
	server.app.Post("/errors", handlers.SaveAppErrors(server.db))

	protected := server.app.Group("/auth", authMiddleware(server.config))

	protected.Post("/parsexls", handlers.ParseXLS())

	protected.Get("/classes/year/:yearid", handlers.ClassListHandler(server.db))
	protected.Post("/class", handlers.NewClassHandler(server.db, server.newID))

	protected.Post("/students", handlers.SaveStudent(server.db))
	protected.Get("/students/year/:yearid", handlers.GetStudents(server.db))

	protected.Post("/classperiods", handlers.SaveClassPeriods(server.db))
	protected.Get("/classperiods/year/:yearid", handlers.GetClassPeriods(server.db))

	protected.Post("/attendancedays", handlers.SaveAttendanceDays(server.db))
	protected.Get("/attendancedays/year/:yearid", handlers.GetAttendanceDays(server.db))

	protected.Post("/attendances", handlers.SaveAttendances(server.db))
	protected.Get("/attendances/year/:yearid", handlers.GetAttendances(server.db))

	protected.Post("/activities", handlers.SaveActivities(server.db))
	protected.Get("/activities/year/:yearid", handlers.GetActivities(server.db))

	protected.Post("/scores", handlers.SaveScores(server.db))
	protected.Get("/scores/year/:yearid", handlers.GetScores(server.db))

	return server.app.Listen(fmt.Sprintf(":%s", server.config.App.Port))
}
