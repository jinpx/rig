package pkg

import (
	"os"
	"path/filepath"

	"rig/pkg/constant"
)

const engineVersion = "v1.0.0"

var (
	appName string
)

func init() {
	if appName == "" {
		appName = os.Getenv(constant.EnvAppName)
		if appName == "" {
			appName = filepath.Base(os.Args[0])
		}
	}
}

func EngineVersion() string {
	return engineVersion
}

// SetName set app anme
func SetName(s string) {
	appName = s
}
