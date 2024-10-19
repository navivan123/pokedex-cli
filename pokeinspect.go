package main

import (
    "fmt"
)

func commandInspect(cfg *cfgCommand) error {
    pokemon, ok := cfg.pokeCaught[cfg.pokeArg]
    if !ok {
        fmt.Println("You have not caught that pokemon, or you have specified an unkown pokemon.  Please check your Pokedex.")
        return nil
    }

    fmt.Println("Name:",   pokemon.Name)
    fmt.Println("Height:", pokemon.Height)
    fmt.Println("Weight:", pokemon.Weight)

    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("   - %v: %v\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, pType := range pokemon.Types {
        fmt.Printf("   - %v\n", pType.Type.Name)
    }
    
    return nil
}
