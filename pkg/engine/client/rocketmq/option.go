package rocketmq

// Config config...
type Config struct {
	Name         string              `json:"name" yaml:"name"`
	Addresses    []string            `json:"addr" yaml:"addr"`
	PushConsumer *PushConsumerConfig `json:"pushConsumer" yaml:"consumer"`
	PullConsumer *PullConsumerConfig `json:"pullConsumer" yaml:"pullConsumer"`
	Producer     *ProducerConfig     `json:"producer" yaml:"producer"`
}
