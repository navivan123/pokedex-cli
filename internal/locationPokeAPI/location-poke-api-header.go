package locationPokeAPI

type LocationPokeApi struct {
    Count    int    `json:"count"`
    Next     string `json:"next,omitempty"`
    Previous string `json:"previous,omitempty"`
    Results  []struct {
    Name string `json:"name"`
    } `json:"results"`
}
