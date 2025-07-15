package usecase

import (
	"github.com/jnates/pokedex/internal/domain/model"
	"github.com/jnates/pokedex/internal/port/input"
	"github.com/jnates/pokedex/internal/port/output"
)

type pokemonUsecase struct {
	repo output.PokemonRepository
}

func NewPokemonUsecase(r output.PokemonRepository) input.GetPokemonUsecase {
	return &pokemonUsecase{repo: r}
}

func (uc *pokemonUsecase) GetPokemon(name string) (*model.Pokemon, error) {
	return uc.repo.GetPokemon(name)
}
