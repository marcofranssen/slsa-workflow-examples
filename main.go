package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	gitVersion = "dev"
	gitCommit  = "unknown"
)

func main() {
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	defer func() {
		_ = logger.Sync()
	}()

	logger.Info("This app was build using SLSA requirements", zap.String("version", gitVersion), zap.String("commit", gitCommit))
}
