package v2

import (
    "fmt"
    "net/http"
)

const DefaultBaseUrl = "https://meta.fabricmc.net/v2"

type Config struct {
    BaseUrl string
    Client  *http.Client
}

// NewConfig returns a default client to use with the API.
func NewConfig() *Config {
    return &Config{
        BaseUrl: DefaultBaseUrl,
        Client:  http.DefaultClient,
    }
}

// GetDownloadUrl returns the URL to download the server jar.
func (c *Config) GetDownloadUrl(gameVersion, loaderVersion, installerVersion string) string {
    url := fmt.Sprintf("%s/versions/loader/%s/%s/%s/server/jar",
        c.BaseUrl,
        gameVersion,
        loaderVersion,
        installerVersion,
    )

    fmt.Printf("Download URL: %s\n", url)
    return url
}
