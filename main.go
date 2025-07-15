package main

import (
	"github.com/jnates/pokedex/internal/infrastructure"
	"github.com/jnates/pokedex/internal/infrastructure/contants"
	"github.com/jnates/pokedex/internal/infrastructure/server"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting API CMD")
	infrastructure.InitLogger()

	port := os.Getenv(contants.API_PORT)
	server.Start(port)
}
