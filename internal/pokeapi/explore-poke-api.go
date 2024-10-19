package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strconv"
)

func (c *Client) CallExploreAPI(area string) (ExplorePokeApi, error) {
    if val, ok := c.cache.Get(area); ok {
        var explore ExplorePokeApi
        if err := json.Unmarshal(val, &explore); err != nil {
            return ExplorePokeApi{}, err
        }

        return explore, nil
    }
    
    url := baseURL + areaURL + area 

    res, err := http.Get(url)
    if err != nil {
        return ExplorePokeApi{}, err
    }
    defer res.Body.Close()

    if res.StatusCode >= http.StatusBadRequest {
        return ExplorePokeApi{}, fmt.Errorf(": %v | %v | Status Code: %v", url, CodeHTTPError, res.StatusCode)
    }

    if res.Header.Get("Content-Type")[0:16] != "application/json" {
        return ExplorePokeApi{}, fmt.Errorf(": %v | %v | Content-Type: %v | Body: %v", url,
                                                                                        BadContentTypeHTTPError,
                                                                                        res.Header.Get("Content-Type"),
                                                                                        res.Body)
    }

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return ExplorePokeApi{}, fmt.Errorf("error reading response: %w", err)
    }

    var explore ExplorePokeApi
    if err = json.Unmarshal(data, &explore); err != nil {
        return ExplorePokeApi{}, err
    }

    c.cache.Add(strconv.Itoa(explore.ID), data)
    c.cache.Add(explore.Name,             data)

    return explore, nil
}
