package initialize
//
//import (
//	"fmt"
//	"github.com/akaedison/go-gin-demo/global"
//	"github.com/akaedison/go-gin-demo/pkg/util"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//	"os"
//	"time"
//)
//
//var level zapcore.Level
//
//func Zap() (logger zap.Logger) {
//	if ok, _ := util.PathExist(global.CFG.Zap.Director); !ok {
//		fmt.Printf("create %v directory\n", global.CFG.Zap.Director)
//		_ = os.Mkdir(global.CFG.Zap.Director, os.ModePerm)
//	}
//
//	switch global.CFG.Zap.Level {
//	case "debug":
//		level = zap.DebugLevel
//	case "info":
//		level = zap.InfoLevel
//	case "warn":
//		level = zap.WarnLevel
//	case "error":
//		level = zap.ErrorLevel
//	case "d_panic":
//		level = zap.DPanicLevel
//	case "panic":
//		level = zap.PanicLevel
//	case "fatal":
//		level = zap.FatalLevel
//	default:
//		level = zap.InfoLevel
//	}
//
//	if level == zap.DebugLevel || level == zap.ErrorLevel {
//
//	}
//
//	return logger
//}
//
//func getEncoderConfig() (config zapcore.EncoderConfig) {
//	config = zapcore.EncoderConfig{
//		MessageKey:       "message",
//		LevelKey:         "level",
//		TimeKey:          "time",
//		NameKey:          "logger",
//		CallerKey:        "caller",
//		StacktraceKey:    global.CFG.Zap.StacktraceKey,
//		LineEnding:       zapcore.DefaultLineEnding,
//		EncodeLevel:      zapcore.LowercaseLevelEncoder,
//		EncodeTime:       CustomTimeEncoder,
//		EncodeDuration:   zapcore.SecondsDurationEncoder,
//		EncodeCaller:     zapcore.FullCallerEncoder,
//	}
//
//	switch {
//	case global.CFG.Zap.EncodeLevel == "LowercaseLevelEncoder":
//		config.EncodeLevel = zapcore.LowercaseLevelEncoder
//	case global.CFG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
//		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
//	case global.CFG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
//		config.EncodeLevel = zapcore.CapitalLevelEncoder
//	case global.CFG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
//		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	default:
//		config.EncodeLevel = zapcore.LowercaseLevelEncoder
//	}
//
//	return config
//}
//
//func getEncoderCore() (core zapcore.Core) {
//}
//
//func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder)  {
//	enc.AppendString(t.Format(global.CFG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
//}