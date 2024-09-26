package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	if cfg.nextLocationAreaURL == nil && cfg.prevLocationAreaURL != nil {
		return errors.New("Cannot move forward any farther")
	}
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("You are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationAreaURL)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	fmt.Println("Location Areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}
