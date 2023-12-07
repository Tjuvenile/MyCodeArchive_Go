package logging

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func init() {
	// json格式的info，默认情况下会打印到console里
	logger, _ := zap.NewProduction()
	// SugaredLogger可以用format的格式记录。Infof，Errorf等。
	Log = logger.Sugar()
}
