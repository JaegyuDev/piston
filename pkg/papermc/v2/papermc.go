package v2

import (
    "fmt"
    "net/http"
)

const (
    DefaultBaseUrl = "https://api.papermc.io/v2"
)

type ProjectId string

const (
    PAPER      ProjectId = "paper"
    TRAVERTINE ProjectId = "travertine"
    WATERFALL  ProjectId = "waterfall"
    VELOCITY   ProjectId = "velocity"
    FOLIA      ProjectId = "folia"
)

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
func (c *Config) GetDownloadUrl(id ProjectId, version string, build int32, download string) string {
    return fmt.Sprintf("%s/projects/%s/versions/%s/builds/%d/downloads/%s",
        c.BaseUrl,
        id,
        version,
        build,
        download,
    )
}
