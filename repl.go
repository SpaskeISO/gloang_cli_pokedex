package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	inputScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		inputScanner.Scan()
		input := inputScanner.Text()
		if inputScanner.Err() != nil {
			log.Fatal("Input scanner ran into an error")
		}

		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		avaliavleCommands := getCommands()

		command, ok := avaliavleCommands[commandName]
		if !ok {
			fmt.Println("Invlaid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next page location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists prevoius page of loaction areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon and add into your pokedex",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all the pokemon in your pokedex",
			callback:    callbackPokedex,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View information about a caught pokemon",
			callback:    callbackInspect,
		},
		"exit": {
			name:        "exit",
			description: "Turns off Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
