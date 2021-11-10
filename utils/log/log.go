package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	logTmFmtWithMS = "2006-01-02 15:04:05.000"
)

var LOG *zap.SugaredLogger

func init() {
	writer := zapcore.AddSync(os.Stdout)
	// 格式相关配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	_log := zap.New(core)
	LOG = _log.Sugar()
	LOG = _log.Sugar()
}
