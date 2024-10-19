package main

import (
    "fmt"
    "math/rand"
    "math"
)

const (
    maxP = 28.2842712475
)

func commandCatch(cfg *cfgCommand) error {

    apiResponse, err := cfg.pokeapiClient.CallCatchAPI(cfg.pokeArg)
    if err != nil {
        return fmt.Errorf("Error while calling Pokemon Location Information API: %v", err)
    }
    
    fmt.Printf("Throwing a Pokeball at %v...\n", cfg.pokeArg)

    randLim := rand.Float64()*maxP

    if randLim < math.Sqrt(float64(apiResponse.BaseXP)) {
        fmt.Printf("%v escaped!\n", apiResponse.Name)
        return nil
    }

    fmt.Printf("%v was caught!\n", apiResponse.Name)

    cfg.pokeCaught[apiResponse.Name] = apiResponse
    return nil

}
