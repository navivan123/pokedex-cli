package locationPokeAPI

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
)

var CodeHTTPError = errors.New("Error: HTTP Response Returned Error-Level Code!")
var BadContentTypeHTTPError = errors.New("Error: Content-Type mismatch!")

func CallLocationAPI(url string) (LocationPokeApi, error) {
    res, err := http.Get(url)
    if err != nil {
        return LocationPokeApi{}, err
    }
    defer res.Body.Close()

    if res.StatusCode >= http.StatusBadRequest {
        return LocationPokeApi{}, fmt.Errorf(": %v | %v | Status Code: %v", url, CodeHTTPError, res.StatusCode)
    }

    if res.Header.Get("Content-Type")[0:16] != "application/json" {
        return LocationPokeApi{}, fmt.Errorf(": %v | %v | Content-Type: %v | Body: %v", url,
                                                                                        BadContentTypeHTTPError,
                                                                                        res.Header.Get("Content-Type"),
                                                                                        res.Body)
    }

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return LocationPokeApi{}, fmt.Errorf("error reading response: %w", err)
    }

    var locations LocationPokeApi
    if err = json.Unmarshal(data, &locations); err != nil {
        return LocationPokeApi{}, err
    }

    return locations, nil
}
