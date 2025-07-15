package server

import (
	"fmt"
	"github.com/jnates/pokedex/internal/adapter/pokeapi"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Start(port string) {
	e := echo.New()
	client := pokeapi.NewClient(5 * time.Second)
	logg

	e.GET("/pokemon/:name", func(c echo.Context) error {
		name := c.Param("name")
		log.Info().Str("pokemon", name).Msg("Solicitud recibida")
		pokemon, err := client.GetPokemon(name)
		if err != nil {
			log.Error().Err(err).Str("pokemon", name).Msg("No se pudo obtener el pokémon")
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, pokemon)
	})

	log.Info().Msgf("Servidor escuchando en puerto %s", port)
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Err(err).Msg("No se pudo iniciar el servidor")
	}
}
