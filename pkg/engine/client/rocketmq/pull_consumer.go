package rocketmq

import (
	"sync/atomic"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type PullConsumer struct {
	rocketmq.PullConsumer
	name string
	PullConsumerConfig

	subscribers  map[string]func()
	interceptors []primitive.Interceptor
	started      *atomic.Bool
	done         chan struct{}
}
