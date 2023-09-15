package class

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	repo := NewClassEntRepo(server.DB, server.DBCtx, server.IDGenerator)
	handlers := NewClassHandler(server.App, server.TeacherApp, repo)
	handlers.handle()
}
