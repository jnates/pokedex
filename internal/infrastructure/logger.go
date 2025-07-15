package infrastructure

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger() {
	_ = godotenv.Load()

	zerolog.TimeFieldFormat = time.RFC3339
	debugMode := parseBool(os.Getenv("LOGGER_DEBUG"))

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	Logger = zerolog.New(output).With().Timestamp().Logger()

	flagDebug := flag.Bool("debug", debugMode, "activar modo debug")
	flag.Parse()

	if *flagDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Logger.Debug().Msg("Modo debug activado")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func parseBool(value string) bool {
	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		Logger.Error().Msgf("Error convirtiendo LOGGER_DEBUG: %v", err)
		return false
	}
	return boolVal
}
