package v2

import "net/http"

const (
    ApiUrl = "https://api.papermc.io/v2"
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
        BaseUrl: "https://api.papermc.io/v2",
        Client:  http.DefaultClient,
    }
}
