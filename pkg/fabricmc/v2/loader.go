package v2

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
)

type LoaderVersion struct {
    Separator string `json:"separator"`
    Build     int    `json:"build"`
    Maven     string `json:"maven"`
    Version   string `json:"version"`
    Stable    bool   `json:"stable"`
}

func (c *Config) GetLoaderVersions() ([]LoaderVersion, error) {
    url := fmt.Sprintf("%s/versions/loader", c.BaseUrl)
    resp, err := c.Client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    var loaderVersions []LoaderVersion
    err = json.Unmarshal(body, &loaderVersions)
    if err != nil {
        return nil, err
    }

    return loaderVersions, nil
}

func (c *Config) GetLatestLoaderVersion() (string, error) {
    loaderVersions, err := c.GetLoaderVersions()
    if err != nil {
        return "", err
    }

    if len(loaderVersions) == 0 {
        return "", errors.New("no loader versions found")
    }

    return loaderVersions[0].Version, nil
}
