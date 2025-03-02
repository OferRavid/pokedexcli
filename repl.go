package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/OferRavid/pokedexcli/internal/pokecache"
)

type locationArea struct {
	Name string `json:"name"`
}

type config struct {
	Results  []locationArea `json:"results"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
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
	}
}

func startRepl(cfg *config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]

		command, exists := getCommands()[commandWord]
		if exists {
			err := command.callback(cfg, cache)
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
