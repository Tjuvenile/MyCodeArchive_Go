package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.SugaredLogger

func init() {
	// 打开文件（如果不存在则创建，以追加写的方式）
	file, err := os.OpenFile("./utils/logging/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("failed to open or create file")
	}

	// 设置日志格式
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "caller",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(caller.TrimmedPath())
		},
	}

	// 配置日志核心
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(file),
		zap.InfoLevel,
	)
	// 默认是会打印到标准输出的，如果只想把日志打印到文件里，可以不打印标准输出，这样单元测试的时候，就不会打印log了
	// witeSyncer = zapcore.NewMultiWriteSyncer(witeSyncer, zapcore.AddSync(os.Stdout))

	//设置为开发模式会记录panic
	development := zap.Development()
	//指定warn和warn之上的级别都需要输出调用堆栈
	enableStacktrace := zap.AddStacktrace(zapcore.WarnLevel)
	//开启记录文件名和行号
	caller := zap.AddCaller()
	// 创建日志器
	logger := zap.New(core, development, enableStacktrace, caller)

	// SugaredLogger 可以用 format 的格式记录，Infof，Errorf 等。
	Log = logger.Sugar()
}
