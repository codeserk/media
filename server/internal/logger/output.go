package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func Output() *zerolog.ConsoleWriter {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006-01-02 15:04:05",
		NoColor:    false,
	}
	output.FormatLevel = func(i interface{}) string {
		level := strings.ToUpper(fmt.Sprintf("%s", i))
		switch level {
		case "INFO":
			return "\u001b[32m[I]\u001b[0m"
		case "WARN":
			return "\u001b[33m[W]\u001b[0m"
		case "ERROR":
			return "\u001b[31m[E]\u001b[0m"
		case "DEBUG":
			return "\u001b[36m[D]\u001b[0m"
		case "FATAL":
			return "\u001b[35m[F]\u001b[0m"
		default:
			return "[" + level + "]"
		}
	}
	output.FormatMessage = func(i interface{}) string {
		if i == nil {
			return "|>"
		}

		return fmt.Sprintf("%s |>", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\u001b[34m%s\u001b[0m=", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	return &output
}
