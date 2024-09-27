package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {

	pokedexEntries := cfg.caughtPokemon
	if len(pokedexEntries) == 0 {
		return fmt.Errorf("No Pokedex entries found")
	}
	fmt.Println("Your Pokedex:")
	for _, entry := range pokedexEntries {
		fmt.Printf("  - %s\n", entry.Name)
	}
	return nil
}
