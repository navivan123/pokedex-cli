package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strconv"
)

func (c *Client) CallCatchAPI(pokemonName string) (CatchPokeApi, error) {
    if val, ok := c.cache.Get(pokemonName); ok {
        var pokemon CatchPokeApi
        if err := json.Unmarshal(val, &pokemon); err != nil {
            return CatchPokeApi{}, err
        }

        return pokemon, nil
    }
    
    url := baseURL + pokeURL + pokemonName 

    res, err := http.Get(url)
    if err != nil {
        return CatchPokeApi{}, err
    }
    defer res.Body.Close()

    if res.StatusCode >= http.StatusBadRequest {
        return CatchPokeApi{}, fmt.Errorf(": %v | %v | Status Code: %v", url, CodeHTTPError, res.StatusCode)
    }

    if res.Header.Get("Content-Type")[0:16] != "application/json" {
        return CatchPokeApi{}, fmt.Errorf(": %v | %v | Content-Type: %v | Body: %v", url,
                                                                                        BadContentTypeHTTPError,
                                                                                        res.Header.Get("Content-Type"),
                                                                                        res.Body)
    }

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return CatchPokeApi{}, fmt.Errorf("error reading response: %w", err)
    }

    var pokemon CatchPokeApi
    if err = json.Unmarshal(data, &pokemon); err != nil {
        return CatchPokeApi{}, err
    }

    c.cache.Add(strconv.Itoa(pokemon.ID), data)
    c.cache.Add(pokemon.Name,             data)

    return pokemon, nil
}
