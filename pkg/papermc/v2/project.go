package v2

import (
    "encoding/json"
    "fmt"
    "io"
)

type ProjectResponse struct {
    // Id PaperMC internal id for project
    Id ProjectId `json:"project_id"`

    // Name User readable name
    Name string `json:"project_name"`

    // VersionGroups the major versions of the game the project supports, i.e. 1.20, 1.21 ...
    VersionGroups []string `json:"version_groups"`

    // Versions the point releases of the game the project supports, i.e. 1.21.1, 1.20.4 ...
    Versions []string `json:"versions"`
}

// GetProject Gets information about a project.
func (c *Config) GetProject(id ProjectId) (ProjectResponse, error) {
    url := fmt.Sprintf("%s/projects/%s", c.BaseUrl, id)
    resp, err := c.Client.Get(url)
    if err != nil {
        return ProjectResponse{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return ProjectResponse{}, err
    }

    var p ProjectResponse
    err = json.Unmarshal(body, &p)
    if err != nil {
        return ProjectResponse{}, err
    }

    return p, nil
}
