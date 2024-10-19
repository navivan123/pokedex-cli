package pokeapi

type CatchPokeApi struct {
    Height int    `json:"height"`
    Weight int    `json:"weight"`
    ID     int    `json:"id"`
    BaseXP int    `json:"base_experience"`
    Name   string `json:"name"`

    Types []struct{
        Type struct {
            Name string `json:"name"`
        } `json"type"`
    } `json:"types"`

    Stats []struct{
        BaseStat int `json:"base_stat"`
        Stat struct {
            Name string `json:"name"`
        } `json:"stat"`
    } `json:"stats"`
}
