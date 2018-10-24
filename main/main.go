package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"zap_rotate_writer/ZapRotateWriter"
)

func main() {
	var err error
	var Logger *zap.Logger
	Logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("Cannot initialize logging")
		return
	}

	ZapRotateSync := &ZapRotateWriter.RotateLogWriteSyncer{}
	ZapRotateSync.RotateLoggerInit("MIDNIGHT", 1, "try.log", 3)
	zapcore.Lock(ZapRotateSync)
	Logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey: "ts",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}), zapcore.NewMultiWriteSyncer(os.Stdout, ZapRotateSync), zap.NewAtomicLevel()))

	defer Logger.Sync()
	for i := 1; i < 100000; i++ {
		Logger.Info("test..." + strconv.Itoa(i))
	}

	Logger.Info("*************************************************** Done ***************************************************")
}
