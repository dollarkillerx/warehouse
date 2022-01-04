package utils

import (
	cfg "github.com/dollarkillerx/common/pkg/config"
	"github.com/dollarkillerx/common/pkg/logger"
)

var Logger *logger.RimeLogger

func InitLogger(config cfg.LoggerConfig) {
	Logger = logger.NewRimeLogger(config)
}
