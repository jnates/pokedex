package output

import "github.com/jnates/pokedex/internal/domain/model"

type PokemonRepository interface {
	GetPokemon(name string) (*model.Pokemon, error)
}
