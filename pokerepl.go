package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type cliCommand struct {
    name            string
    description     string
    callback        func(*cfgCommand) (error)
}

type cfgCommand struct {
    pokeUrlNext     string
    pokeUrlPrevious string
}

func pokeRepl() {
    comHelp := cliCommand{ name: "help", description: "Displays a Help Message", callback: commandHelp, }
    comExit := cliCommand{ name: "exit", description: "Exit the Pokedex", callback: commandExit, }
    comMap  := cliCommand{ name:        "map", 
                           description: "View Next 20 Locations in the Pokemon World!", 
                           callback:    commandMap, }
    comMapb := cliCommand{ name: "mapb", description: "View Previous 20 Locations in the Pokemon World!", callback: commandMapb, }
    
    configs := map[string]*cliCommand{ "help": &comHelp, "exit": &comExit, "map": &comMap, "mapb": &comMapb }
    comCfg  := cfgCommand{ pokeUrlNext:     "https://pokeapi.co/api/v2/location-area/", 
                           pokeUrlPrevious: "https://pokeapi.co/api/v2/location-area/" }

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("PokeCLI > ")
        scanner.Scan()

        lowerCaseWords := strings.ToLower(scanner.Text())
        words := strings.Fields(lowerCaseWords)
        if len(words) == 0 {
            continue
        }

        commandName := words[0]

        command, ok := configs[commandName]
        if ok {
            err := command.callback(&comCfg)
            if err != nil {
                fmt.Println(err)
            }

        } else {
            fmt.Println("Unknown Command!")
        }
    }
}


