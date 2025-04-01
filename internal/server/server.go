package server

import "rig/internal/configuration"

func Biz() *engine_gin.ServerEngine {
	var (
		server = configuration.Server.ServerHttp.MustBuild()
	)
	return server
}
