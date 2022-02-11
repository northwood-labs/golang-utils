package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// GetStdTextLogger is a standard plain-text logger.
func GetStdTextLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	return log
}

// GetStdJSONLogger is a standard plain-text logger.
func GetStdJSONLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFormatUnix,
	}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	return log
}
