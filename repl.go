package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OferRavid/pokedexcli/internal/pokeapi"
)

type config struct {
	caughtPokemon        map[string]pokeapi.Pokemon
	caughtPokemonCount   map[string]int
	areaExplored         []string
	pokeapiClient        pokeapi.Client
	NextLocationsURL     *string `json:"next"`
	PreviousLocationsURL *string `json:"previous"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

/*
getCommands returns a map of available CLI commands, each associated with a name,
description, and a corresponding callback function.
*/
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next batch of 20 location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous batch of 20 location-areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays all Pokemon in the area given",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a Pokemon encountered in an area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Shows details about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the Pokemon you caught",
			callback:    commandPokedex,
		},
	}
}

/*
startRepl initializes and runs the REPL (Read-Eval-Print Loop) for the Pokedex CLI.
It continuously reads user input, processes commands, and executes corresponding functions.
*/
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

/*
cleanInput processes the user input by converting it to lowercase,
trimming spaces, and splitting it into individual words.
Returns a slice of cleaned words.
*/
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)
	return strings.Fields(trimmed)
}
