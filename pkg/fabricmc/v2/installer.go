package v2

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
)

type InstallerVersion struct {
    Url     string `json:"url"`
    Maven   string `json:"maven"`
    Version string `json:"version"`
    Stable  bool   `json:"stable"`
}

func (c *Config) GetInstallerVersions() ([]InstallerVersion, error) {
    url := fmt.Sprintf("%s/versions/installer", c.BaseUrl)
    resp, err := c.Client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var installerVersions []InstallerVersion
    err = json.Unmarshal(body, &installerVersions)
    if err != nil {
        return nil, err
    }

    return installerVersions, nil
}

func (c *Config) GetLatestInstallerVersion() (string, error) {
    // For now we're going to assume a stable installer version supports snapshots. This needs testing.
    installerVersions, err := c.GetInstallerVersions()
    if err != nil {
        return "", err
    }

    if len(installerVersions) == 0 {
        return "", errors.New("no installer versions found")
    }

    return installerVersions[0].Version, nil
}
