package main

import (
    "fmt"
    "internal/locationPokeAPI"
)

func commandMap(cfg *cfgCommand) error {
    err := commandMapHelper(cfg, cfg.pokeUrlNext)
    if err != nil {
        return fmt.Errorf("%v | End of list!", err)
    }
    return nil
}

func commandMapb(cfg *cfgCommand) error {
    err := commandMapHelper(cfg, cfg.pokeUrlPrevious)
    if err != nil {
        return fmt.Errorf("%v | Start of list!", err)
    }
    return nil
}

func commandMapHelper(cfg *cfgCommand, pokeUrl string) error {
    if pokeUrl == "" {
        return fmt.Errorf("Error while calling Pokemon Location API: No URL!")
    }

    apiResponse, err := locationPokeAPI.CallLocationAPI(pokeUrl)
    if err != nil {
        return fmt.Errorf("Error while calling Pokemon Location API: %v", err)
    }

    if apiResponse.Next == "" && apiResponse.Previous != "" {
        fmt.Println("At the end of location page!")
        cfg.pokeUrlNext     = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=%v", 
                                          apiResponse.Count - apiResponse.Count % 20, 
                                          apiResponse.Count % 20)
        cfg.pokeUrlPrevious = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", 
                                          apiResponse.Count - (apiResponse.Count % 20) - 20)

    } else if apiResponse.Next != "" && apiResponse.Previous == "" {
        fmt.Println("At the start of location page!")
        cfg.pokeUrlPrevious = "https://pokeapi.co/api/v2/location-area/"
        cfg.pokeUrlNext = "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"

    } else if apiResponse.Next == "" && apiResponse.Previous == "" {
        cfg.pokeUrlPrevious = "https://pokeapi.co/api/v2/location-area/"
        cfg.pokeUrlNext = "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"
        return fmt.Errorf("Unknown state of location config. Resetting to defaults!")

    } else {
        cfg.pokeUrlNext     = apiResponse.Next
        cfg.pokeUrlPrevious = apiResponse.Previous
    }

    for _, location := range apiResponse.Results {
        fmt.Printf("%v\n", location.Name)
    }

    return nil
}
