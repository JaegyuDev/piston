package v2

import (
    "encoding/json"
    "fmt"
    "io"
)

type ProjectsResponse struct {
    Projects []string `json:"projects"`
}

// GetProjects Gets a list of all available projects.
func (c *Config) GetProjects() (ProjectsResponse, error) {
    url := fmt.Sprintf("%s/projects", c.BaseUrl)
    resp, err := c.Client.Get(url)
    if err != nil {
        return ProjectsResponse{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return ProjectsResponse{}, err
    }

    var p ProjectsResponse
    err = json.Unmarshal(body, &p)
    if err != nil {
        return ProjectsResponse{}, err
    }

    return p, nil
}
