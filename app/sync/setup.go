package sync

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	repo := NewSyncEntRepo(server.DB, server.DBCtx)
	handler := NewSyncHandler(server.TeacherApp, repo)
	handler.handle()
}
