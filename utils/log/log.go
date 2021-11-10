package log

import (
	"fmt"
	"go.uber.org/zap"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

var LOG *zap.SugaredLogger

func init() {
	_log, err := zap.NewProduction()
	if err != nil {
		_ = fmt.Errorf("Init ZapLogger failed, err: %v ", err)
	}
	LOG = _log.Sugar()
}
