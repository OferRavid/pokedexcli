package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PreviousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.PreviousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.PreviousLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PreviousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	cfg.areaExplored = []string{}
	cfg.areaExplored = append(cfg.areaExplored, location.Name)
	for _, enc := range location.PokemonEncounters {
		name := enc.Pokemon.Name
		fmt.Printf(" - %s\n", name)
		cfg.areaExplored = append(cfg.areaExplored, name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	if len(cfg.areaExplored) == 0 {
		return errors.New("you must explore an area for Pokemon encounters first")
	}
	if ok := slices.Contains(cfg.areaExplored, name); !ok {
		return fmt.Errorf("you didn't encounter %s in %s.\nexplore %s again to see the Pokemon encountered", name, cfg.areaExplored[0], cfg.areaExplored[0])
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if success := rand.Intn(pokemon.BaseExperience) <= 40; success {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		if _, ok := cfg.caughtPokemonCount[pokemon.Name]; ok {
			cfg.caughtPokemonCount[pokemon.Name]++
		} else {
			cfg.caughtPokemonCount[pokemon.Name] = 1
			cfg.caughtPokemon[pokemon.Name] = pokemon
		}
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Printf("  - %s\n", typeInfo.Type.Name)
		}
		return nil
	}

	return fmt.Errorf("can't show information on %s. you need to catch one first", name)
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) > 0 {
		fmt.Println("Your Pokedex:")
		for name, _ := range cfg.caughtPokemon {
			fmt.Printf(" - %s\n", name)
		}
		return nil
	}

	return errors.New("your pokedex is empty. go catch some pokemon")
}
