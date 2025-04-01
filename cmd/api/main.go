package main

import (
	"rig/internal/configuration"
	"rig/internal/server"

	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg"
)

func main() {
	var (
		eng = &app.ServerEngine{}
	)

	pkg.SetName("game")

	if err := eng.Startup(
		configuration.GetConfig,
	); err != nil {
		panic(err)
	}

	if err := eng.Run(
		server.Biz(),
	); err != nil {
		panic(err)
	}
}
