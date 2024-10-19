package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "internal/pokeapi"
)

type cliCommand struct {
    name            string
    description     string
    callback        func(*cfgCommand) (error)
}

type cfgCommand struct {
    pokeUrlNext     string
    pokeUrlPrevious string
    pokeArg         string
    pokeCaught      map[string]pokeapi.CatchPokeApi
    pokeapiClient   pokeapi.Client
}

func pokeRepl(comCfg *cfgCommand) {
    comHelp    := cliCommand{ name: "help",    description: "Displays a Help Message",                                   callback: commandHelp,    }
    comExit    := cliCommand{ name: "exit",    description: "Exit the Pokedex",                                          callback: commandExit,    }
    comMap     := cliCommand{ name: "map",     description: "View Next 20 Locations in the Pokemon World!",              callback: commandMap,     }
    comMapb    := cliCommand{ name: "mapb",    description: "View Previous 20 Locations in the Pokemon World!",          callback: commandMapb,    }
    comExplore := cliCommand{ name: "explore", description: "View Pokemons Available to Capture in the Area Specified!", callback: commandExplore, }
    comCatch   := cliCommand{ name: "catch",   description: "Attempt to Catch Specified Pokemon!",                       callback: commandCatch,   }
    comInspect := cliCommand{ name: "inspect", description: "Inspect Details About a Specified Pokemon!",                callback: commandInspect, }
    comPokedex := cliCommand{ name: "pokedex", description: "View All Pokemon You Have Caught!",                         callback: commandPokedex, }

    configs := map[string]*cliCommand{ "help":    &comHelp,    "exit":  &comExit,  "map": &comMap,         "mapb":    &comMapb, 
                                       "explore": &comExplore, "catch": &comCatch, "inspect": &comInspect, "pokedex": &comPokedex, }

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
            if command.name == "explore" || command.name == "catch" || command.name == "inspect" {
                if len(words) < 2 {
                    fmt.Println("Error, please provide area name or id when using explore command.")
                    continue
                }
                comCfg.pokeArg = words[1]
            } 

            err := command.callback(comCfg)
            if err != nil {
                fmt.Println(err)
            }

        } else {
            fmt.Println("Unknown Command!")
            command, _ = configs["help"]
            command.callback(comCfg)
        }
    }
}
