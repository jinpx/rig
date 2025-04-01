package rocketmq

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/juju/ratelimit"
)

type PushConsumer struct {
	rocketmq.PushConsumer
	name string
	PushConsumerConfig

	subscribers  map[string]func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)
	interceptors []primitive.Interceptor
	bucket       *ratelimit.Bucket
	started      bool
}
