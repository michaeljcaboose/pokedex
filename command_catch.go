package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	if randNum > threshold {
		return fmt.Errorf("%s escaped!", name)
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.caughtPokemon[name] = pokemon

	return nil
}
