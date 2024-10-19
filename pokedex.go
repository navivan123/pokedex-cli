package main

import (
    "fmt"
)

func commandPokedex(cfg *cfgCommand) error {

    if len(cfg.pokeCaught) <= 0 {
        fmt.Println("You have not caught any pokemon! Use the 'catch' command and specify the pokemon you want to catch.")
        return nil
    }

    fmt.Println("Your Pokedex:")

    for pokemon, _ := range cfg.pokeCaught {
        fmt.Println("   -", pokemon)
    }

    return nil
}
