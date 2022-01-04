package config

import (
	cfg "github.com/dollarkillerx/common/pkg/config"

	"os"
	"strings"
)

func GetLoggerConfig() cfg.LoggerConfig {
	logPath := os.Getenv("LogPath")
	if logPath == "" {
		logPath = "/log/warehouse.log"
	}

	return cfg.LoggerConfig{
		Filename:   logPath,
		MaxSize:    200,
		MaxBackups: 3,
	}
}

func GetCORSAllowedOrigins() []string {
	cors := os.Getenv("CORSAllowedOrigins")
	if cors == "" {
		return nil
	}

	return strings.Split(cors, ",")
}

func GetListenAddr() string {
	addr := os.Getenv("ListenAddr")
	if addr == "" {
		return ":8187"
	}

	return addr
}

func GetAccessKey() string {
	return os.Getenv("AccessKey")
}

func GetSecretKey() string {
	return os.Getenv("SecretKey")
}
