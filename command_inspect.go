package main

import (
	"errors"
	"fmt"

	"math/rand"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught %s yet", pokemon.Name)
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if randNum > 50 {
		return fmt.Errorf("%s broke out of the pokeball", pokemon.Name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Pokemon Type(s):")
	for i, typ := range pokemon.Types {
		fmt.Printf(" - Type %v: %s\n", i+1, typ.Type.Name)
	}
	fmt.Println("Base Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
