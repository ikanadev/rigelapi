package extra

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	repo := NewExtraEntRepo(server.DB, server.DBCtx)
	handlers := NewExtraHandler(server.App, server.TeacherApp, repo)
	handlers.handle()
}
