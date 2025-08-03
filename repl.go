package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	// Beginning of commands section
	commands := make(map[string]cliCommand)
	cfg := &config{}
	cfg.Next = "https://pokeapi.co/api/v2/location-area/"

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback: func() error {
			return commandExit(cfg)
		},
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback: func() error {
			return commandHelp(cfg, commands)
		},
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 location-areas",
		callback: func() error {
			return commandMap(cfg)
		},
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location-areas",
		callback: func() error {
			return commandMapb(cfg)
		},
	}

	// Main REPL Loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		commandSlice := cleanInput(command)

		commandStruct := commands[commandSlice[0]]
		if commandStruct.callback == nil {
			fmt.Println("Unknown command")
			continue
		}
		if commandStruct.callback != nil {
			commandStruct.callback()
		}
	}
}

func cleanInput(text string) []string {

	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, commands map[string]cliCommand) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, key := range commands {
		fmt.Println(key.name+": ", key.description)
	}

	return nil
}

type config struct {
	Next     string
	Previous string
}

type locationArea struct {
	Name string `json:"name"`
}

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

func commandMap(c *config) error {
	// Add the map command. It displays the names of 20 location areas in the

	res, err := http.Get(c.Next)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Println(err)
	}

	resp := locationAreaResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println(err)
	}

	c.Next = resp.Next
	c.Previous = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(c *config) error {

	// Do the same as commandMap, but in reverse

	// Make sure there is a previous page
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(c.Previous)
	if err != nil {
		log.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Println(err)
	}

	resp := locationAreaResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println(err)
	}

	c.Next = resp.Next
	c.Previous = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
