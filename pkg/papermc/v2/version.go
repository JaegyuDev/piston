package v2

import (
    "encoding/json"
    "fmt"
    "io"
)

type VersionResponse struct {
    Id      ProjectId `json:"project_id"`
    Name    string    `json:"project_name"`
    Version string    `json:"version"`
    Builds  []int32   `json:"builds"`
}

// GetVersion Gets information about a version.
func (c *Config) GetVersion(id ProjectId, version string) (VersionResponse, error) {
    url := fmt.Sprintf("%s/projects/%s/versions/%s", c.BaseUrl, id, version)
    resp, err := c.Client.Get(url)
    if err != nil {
        return VersionResponse{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return VersionResponse{}, err
    }

    var v VersionResponse
    err = json.Unmarshal(body, &v)
    if err != nil {
        return VersionResponse{}, err
    }

    return v, nil
}
