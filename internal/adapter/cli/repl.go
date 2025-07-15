package cli

import (
	"bufio"
	"fmt"
	"github.com/jnates/pokedex/internal/port/input"
	"os"
	"strings"
)

func StartREPL(usecase input.GetPokemonUsecase) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command > ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		if text == "exit" {
			fmt.Println("Bye!")
			break
		}
		pokemon, err := usecase.GetPokemon(text)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Name: %s | Height: %d | Weight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	}
}
