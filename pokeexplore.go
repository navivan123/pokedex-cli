package main

import (
    "fmt"
)

func commandExplore(cfg *cfgCommand) error {
    if cfg.pokeArg == "" {
        return fmt.Errorf("Area not valid! Please select an area.")
    }
    
    apiResponse, err := cfg.pokeapiClient.CallExploreAPI(cfg.pokeArg)
    if err != nil {
        return fmt.Errorf("Error while calling Pokemon Location Information API: %v", err)
    }

    for _, pokemon := range apiResponse.PokemonEncounters {
        fmt.Printf("%v\n", pokemon.Pokemon.Name)
    }

    return nil
}
