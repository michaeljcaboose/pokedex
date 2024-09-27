package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("No pokemon provided")
	}

	name := args[0]

	pokemon, ok := cfg.caughtPokemon[name]
	if ok != true {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, poketype := range pokemon.Types {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}

	return nil
}
