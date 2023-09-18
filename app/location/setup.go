package location

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	entRepo := NewLocationEntRepo(server.DB, server.DBCtx)
	handlers := NewLocationHandler(server.App, entRepo)
	handlers.handle()
}
