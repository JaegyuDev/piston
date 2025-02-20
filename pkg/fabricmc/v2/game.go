package v2

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
)

type GameVersion struct {
    Version string `json:"version"`
    Stable  bool   `json:"stable"`
}

func (c *Config) GetGameVersions() ([]GameVersion, error) {
    url := fmt.Sprintf("%s/versions/game", c.BaseUrl)
    resp, err := c.Client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    var gameVersions []GameVersion
    err = json.Unmarshal(body, &gameVersions)
    if err != nil {
        return nil, err
    }

    return gameVersions, nil

}

func (c *Config) GetGameVersion(v string) (bool, GameVersion) {
    versions, err := c.GetGameVersions()
    if err != nil {
        // tbh this should always work unless net conn not working, I just want this method to have a nice interface
        // if the user conn is spotty we'll shoot them irl
        fmt.Printf("Error getting versions: %s\n", err)
        return false, GameVersion{}
    }

    for _, version := range versions {
        if version.Version == v {
            return true, version
        }
    }

    return false, GameVersion{}
}

func (c *Config) GetLatestGameVersion(snapshot bool) (string, error) {
    versions, err := c.GetGameVersions()
    if err != nil {
        return "", err
    }

    for _, version := range versions {
        if !snapshot && version.Stable {
            return version.Version, nil
        }

        return version.Version, nil
    }

    return "", errors.New("no versions found")
}
