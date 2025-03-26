package u_go

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"rig/pkg/utils/u_string"

	"gitlab.casinovip.tech/minigame_backend/c_engine/pkg/log"
)

var (
	_logger = log.Engine().With(zap.String("mod", "u_go")).WithOptions(zap.AddStacktrace(zap.ErrorLevel))
)

func try(fn func() error, cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			_, file, line, _ := runtime.Caller(2)

			_logger.Error("recover", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
			ret = errors.Wrap(ret, fmt.Sprintf("%s:%d", u_string.FunctionName(fn), line))
		}
	}()
	return fn()
}

func safe_try(fn func(), cleaner func()) (ret error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		if err := recover(); err != nil {
			_, file, line, _ := runtime.Caller(2)

			_logger.Error("recover", zap.Any("err", err), zap.String("line", fmt.Sprintf("%s:%d", file, line)))
			if _, ok := err.(error); ok {
				ret = err.(error)
			} else {
				ret = fmt.Errorf("%+v", err)
			}
			ret = errors.Wrap(ret, fmt.Sprintf("%s:%d", u_string.FunctionName(fn), line))
		}
	}()
	fn()
	return nil
}
