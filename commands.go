package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
)

/*
commandHelp prints a help message displaying available commands and their descriptions.
It takes a config struct but does not use it.
*/
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

/*
commandExit prints a farewell message and exits the program.
It does not return an error.
*/
func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

/*
commandMap retrieves and prints the next batch of location areas using the PokeAPI client.
It updates the configuration with new pagination URLs.
*/
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

/*
commandMapb retrieves and prints the previous batch of location areas using the PokeAPI client.
It updates the configuration with new pagination URLs.
*/
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

/*
commandExplore retrieves and displays the Pokemon encountered in a specified location.
It updates the explored area in the configuration.
*/
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

/*
commandCatch attempts to catch a specified Pokemon encountered in an explored area.
It checks if the Pokemon was encountered before allowing the capture attempt.
*/
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

/*
commandInspect displays detailed information about a caught Pokemon.
If the Pokemon has not been caught, it returns an error.
*/
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Printf(" - %s\n", typeInfo.Type.Name)
		}
		return nil
	}

	return fmt.Errorf("can't show information on %s. you need to catch one first", name)
}

/*
commandPokedex displays a list of all caught Pokemon.
If no Pokemon have been caught, it returns an error.
*/
func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) > 0 {
		fmt.Println("Your Pokedex:")
		for name := range cfg.caughtPokemon {
			fmt.Printf(" - %s\n", name)
		}
		return nil
	}

	return errors.New("your pokedex is empty. go catch some pokemon")
}
