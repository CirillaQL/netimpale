package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var LOG *zap.SugaredLogger

// TimeEncoder 自定义日志的时间输出格式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// init log包初始化
func init() {
	writer := zapcore.AddSync(os.Stdout)
	// 格式相关配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	_log := zap.New(core)
	LOG = _log.Sugar()
}
