package pokeapi

import (
	"encoding/json"
	"fmt"
	structPoke "github.com/jnates/pokedex/internal/adapter/model"
	"github.com/jnates/pokedex/internal/domain/model"
	"github.com/jnates/pokedex/internal/port/output"
	"net/http"
	"time"
)

type pokeAPIClient struct {
	client *http.Client
}

func NewClient(timeout time.Duration) output.PokemonRepository {
	return &pokeAPIClient{
		client: &http.Client{Timeout: timeout},
	}
}

func (p *pokeAPIClient) GetPokemon(name string) (*model.Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw structPoke.Raw
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	return &model.Pokemon{
		Name:   raw.Name,
		Height: raw.Height,
		Weight: raw.Weight,
	}, nil
}
