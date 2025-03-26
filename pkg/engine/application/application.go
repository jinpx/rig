package application

import (
	"flag"
	"log"
	"rig/pkg/utils/u_go"
	"sync"

	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg/core/component"
	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg/utils/u_cycle"
)

type TApplication struct {
	smu        *sync.RWMutex
	initOnce   sync.Once
	stopOnce   sync.Once
	logger     *log.Logger
	disableMap map[Disable]bool
	HideBanner bool
	stopped    chan struct{}
}

// initialize application
func (app *TApplication) initialize() {

	app.initOnce.Do(func() {
		//assign
		app.cycle = u_cycle.NewCycle()
		app.smu = &sync.RWMutex{}
		app.logger = log.Engine()
		app.disableMap = make(map[Disable]bool)
		app.stopped = make(chan struct{})
		app.components = make([]component.Component, 0)
		//private method

		_ = app.parseFlags()
		_ = app.printBanner()
		// app.initLogger()
	})
}

func (app *TApplication) Startup(fns ...func() error) error {

	return u_go.SerialUntilError(fns...)()
}

// parseFlags init
func (app *TApplication) parseFlags() error {
	if app.isDisable(DisableParserFlag) {
		// app.logger.Info("parseFlags disable", log.FieldMod(code.ModApp))
		return nil
	}

	return flag.Parse()
}
