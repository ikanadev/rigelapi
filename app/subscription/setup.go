package subscription

import "github.com/vmkevv/rigelapi/app"

func Setup(server app.Server) {
	repo := NewSubscriptionEntRepo(server.DB, server.DBCtx, server.IDGenerator)
	handler := NewSubscriptionHandler(server.AdminApp, repo)
	handler.handle()
}
