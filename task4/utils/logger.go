package utils

import (
	"io"
	"os"
	"time"

	"task4/config"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger() error {
	cfg := config.AppConfig.Logging

	Logger = logrus.New()

	switch cfg.Level {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

	if cfg.Format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	var output io.Writer
	if cfg.Output == "" {
		output = os.Stdout
	} else {
		file, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		output = io.MultiWriter(os.Stdout, file)
	}

	Logger.SetOutput(output)
	return nil
}
