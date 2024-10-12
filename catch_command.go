package main

import (
	"errors"
	"fmt"

	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if randNum > 50 {
		return fmt.Errorf("%s broke out of the pokeball", pokemon.Name)
	}

	fmt.Printf("%s was caught!", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
