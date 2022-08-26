package logs

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	logfile, _ := os.OpenFile("./zap.log", os.O_CREATE|os.O_APPEND, 0666)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(logfile), zapcore.InfoLevel)

	lg := zap.New(core, zap.AddCaller())
	logger := lg.Sugar()
	logger.Info("info 日志", zap.String("参数 k", "参数 v"))
}

func TestLog2(t *testing.T) {
	//	使用 Lumberjack 进行日志切割归档,zap 本身不切割日志

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,     // M
		MaxBackups: 2,     // 最大数量
		MaxAge:     30,    // 旧文件保存天数
		Compress:   false, // false  不归档压缩旧文件
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(lumberJackLogger), zapcore.InfoLevel)
	lg := zap.New(core, zap.AddCaller())
	for {
		lg.Info("info 日志", zap.String("参数 k", "参数 v"))
	}
}
