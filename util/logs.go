package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logs = &zap.Logger{}

//
func Logs() *zap.SugaredLogger {
	return zap.S()
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.DebugLevel
	}
}

func InitLogs(logPath, level string) (err error) {
	if err = CheckIfMakeDir(logPath); err != nil {
		return err
	}

	logFile := path.Join(logPath, "flowpipe.log")

	cfg := zap.Config{
		Encoding:          "console", //json
		Level:             zap.NewAtomicLevelAt(getLevel(level)),
		Development:       false,
		DisableStacktrace: true,
		OutputPaths:       []string{"stderr", logFile},
		ErrorOutputPaths:  []string{"stderr", logFile},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			StacktraceKey: "stacktrace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			TimeKey:       "time",
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			CallerKey:     "caller",
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	logs, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	zap.RedirectStdLog(logs)
	zap.ReplaceGlobals(logs)
	return
}

// 判断日志文件是否存在并创建
func CheckIfMakeDir(logPath string) error {
	_, err := os.Stat(logPath)
	if err != nil {
		fmt.Println("woshicanwei")
		if os.IsNotExist(err) {
			mask := syscall.Umask(0)
			defer syscall.Umask(mask)
			// 创建日志目录
			fmt.Println("woshicanwei")
			err := os.MkdirAll(filepath.Join(logPath, "log"), 0766)
			if err != nil {
				fmt.Errorf("%s", "创建日志文件失败")
				return err
			}
		}
		return err
	}
	return err
}
