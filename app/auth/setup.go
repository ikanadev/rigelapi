package auth

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	entRepo := NewAuthEntRepo(server.DB, server.DBCtx, server.Config, server.IDGenerator)
	handlers := NewAuthHandler(server.App, server.TeacherApp, server.AdminApp, entRepo, server.Config)
	handlers.handle()
}
