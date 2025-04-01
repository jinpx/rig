package free

import "time"

type Config struct {
	Size         Size          // 缓存容量,最小512*1024 【必填】
	Expire       time.Duration // 失效时间 【必填】
	EnableMetric bool          // metric上报
	Name         string        // 本地缓存名称，用于日志标识&metric上报【选填】
}
