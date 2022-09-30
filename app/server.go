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
	server.app.Use(logUserAgent())

	server.app.Post("/signup", handlers.SignUpHandler(server.db, server.newID))
	server.app.Post("/signin", handlers.SignInHandler(server.db, server.config))
	server.app.Get("/deps", handlers.DepsHandler(server.db))
	server.app.Get("/provs/dep/:depid", handlers.ProvsHandler(server.db))
	server.app.Get("/muns/prov/:provid", handlers.MunsHandler(server.db))
	server.app.Get("/schools/mun/:munid", handlers.SchoolsHandler(server.db))
	server.app.Get("/years", handlers.YearlyDataHandler(server.db))
	server.app.Get("/static", handlers.StaticDataHandler(server.db))

	protected := server.app.Group("/auth", authMiddleware(server.config))

	protected.Get("/classes/year/:yearid", handlers.ClassListHandler(server.db))
	protected.Post("/class", handlers.NewClassHandler(server.db, server.newID))

	protected.Post("/students", handlers.SaveStudent(server.db, server.newID))
	protected.Get("/student/sync", handlers.StudentSyncStatus(server.db))
	protected.Get("/students/year/:yearid", handlers.GetStudents(server.db))

	protected.Post("/classperiods", handlers.SaveClassPeriods(server.db, server.newID))
	protected.Get("/classperiod/sync", handlers.ClassPeriodSyncStatus(server.db))
	protected.Get("/classperiods/year/:yearid", handlers.GetClassPeriods(server.db))

	protected.Post("/attendancedays", handlers.SaveAttendanceDays(server.db, server.newID))
	protected.Get("/attendanceday/sync", handlers.AttendanceDaySyncStatus(server.db))
	protected.Get("/attendancedays/year/:yearid", handlers.GetAttendanceDays(server.db))

	server.app.Listen(fmt.Sprintf(":%s", server.config.App.Port))
}
