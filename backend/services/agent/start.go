package main

import (
	"github.com/becker63/simtricky/common/logger"
	"go.uber.org/zap"
)

func start() {
	log := logger.GetLogger()
	defer log.Sync()

	log.Info("Starting agent service",
		zap.String("service", "agent"),
		zap.Int("port", 8081),
	)
}
