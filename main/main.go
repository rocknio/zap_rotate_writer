package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"time"
	"zap_rotate_writer/ZapRotateWriter"
)

func main() {
	//var err error
	//var Logger *zap.Logger
	//Logger, err = zap.NewProduction()
	//if err != nil {
	//	fmt.Println("Cannot initialize logging")
	//	return
	//}
	//Logger.Info("test, test")

	ZapRotateSync := &ZapRotateWriter.RotateLogWriteSyncer{}
	ZapRotateSync.RotateLoggerInit("MIDNIGHT", 0, "try.log", 3)
	zapcore.Lock(ZapRotateSync)

	productionEncoderConfig := zap.NewProductionEncoderConfig()
	productionEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	Logger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(productionEncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, ZapRotateSync), zap.NewAtomicLevelAt(zapcore.InfoLevel)), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	defer Logger.Sync()
	for i := 1; i < 100000; i++ {
		Logger.Info("test >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> " + strconv.FormatInt(time.Now().Unix(), 10))

		Logger.Error("Kartor HandleKafkaMessage", zap.String("topic", "topic"), zap.String("key", "key"), zap.String("value", "value"))

		Logger.Debug("test >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> " + strconv.FormatInt(time.Now().Unix(), 10))
	}

	Logger.Info("*************************************************** Done ***************************************************")
}
