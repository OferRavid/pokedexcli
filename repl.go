package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OferRavid/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	NextLocationsURL     *string `json:"next"`
	PreviousLocationsURL *string `json:"previous"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
			description: "Explore a location",
			callback:    commandExplore,
		},
	}
}

func startRepl(cfg *config, args ...string) {
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

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)
	return strings.Fields(trimmed)
}
