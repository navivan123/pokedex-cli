package main

import(
    "time"
    "internal/pokeapi"
)



func main() {
    pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
    comCfg  := &cfgCommand{ pokeUrlNext:     "https://pokeapi.co/api/v2/location-area/", 
                            pokeUrlPrevious: "https://pokeapi.co/api/v2/location-area/",
                            pokeapiClient: pokeClient, 
                            pokeArg: "", 
                            pokeCaught: map[string]pokeapi.CatchPokeApi{},
                          }
    pokeRepl(comCfg)
}
