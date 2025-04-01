package data

import (
	"rig/pkg/engine/client/es"
	"rig/pkg/engine/client/redis"
	"rig/pkg/engine/client/rocketmq"
	"rig/pkg/engine/store/cache/free"

	"gorm.io/gorm"
)

type CommonConfig struct {
	MySQL    gorm.Config
	Redis    redis.Config
	Elastic  es.Config
	Rocket   rocketmq.Config
	Cache    free.Config
	TidbConf gorm.Config
}

type ServerConfig struct {
	ServerHttp *engine_gin.Config  `json:"serverHttp" yaml:"serverHttp"`
	ServerRpc  *engine_grpc.Config `json:"serverRpc" yaml:"serverRpc"`
}
