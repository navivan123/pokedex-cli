package main

import "fmt"

func commandHelp(_ *cfgCommand) error {
    fmt.Println()
    fmt.Println("Welcome to PokeCLI! Use commands to grab Pokemon data.")
    fmt.Println()
    fmt.Println("Usage:")

    for _, com := range getCommands() {
        fmt.Printf("      %v: %v\n", com.name, com.description)
    }
    fmt.Println()
    return nil
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a Help Message",
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
        },
        "map": {
            name:        "map",
            description: "View Next 20 Locations in the Pokemon World!",
        },
        "mapb": {
            name:        "mapb",
            description: "View Previous 20 Locations in the Pokemon World!",
        },
    }
}
