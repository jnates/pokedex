package input

import "github.com/jnates/pokedex/internal/domain/model"

type GetPokemonUsecase interface {
	GetPokemon(name string) (*model.Pokemon, error)
}
