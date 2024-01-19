package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Logger() zerolog.Logger {
	// logger file setup
	file, err := os.OpenFile("file.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}
	logger := zerolog.New(file).With().Timestamp().Logger()
	return logger
}
