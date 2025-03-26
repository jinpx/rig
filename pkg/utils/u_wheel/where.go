package u_wheel

import (
	"context"
	"rig/pkg/utils/u_go"
	"time"
)

type THandler interface {
	ProcessTask(ctx context.Context, task *TTask) error
}

// TTask - 任务结构
type TTask struct {
	pattern  string
	payload  []byte
	duration time.Duration
}

func (t *TTask) Payload() []byte {
	return t.payload
}

func NewTask(pattern string, duration time.Duration, payload []byte) *TTask {
	return &TTask{
		pattern:  pattern,
		duration: duration,
		payload:  payload,
	}
}

// TCycle - 循环结构
type TCycle struct {
	tick    int64 // in milliseconds
	collect *TCollection
}

func NewCycle() *TCycle {
	return &TCycle{
		tick:    1,
		collect: NewCollection(),
	}
}

func (m *TCycle) After(ctx context.Context, task *TTask) {
	u_go.Go(func() {
		if task.duration > 0 {
			time.Sleep(task.duration)
		}
		handler, ok := m.collect.Get(task.pattern)
		if ok {
			_ = handler.(THandler).ProcessTask(ctx, task)
		} else {
			// log.Error("wheel.cycle After ? ?", task.pattern)
		}
	})
}

func (m *TCycle) Handle(pattern string, handler THandler) {
	m.collect.Set(pattern, handler)
}
