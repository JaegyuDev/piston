package v2

import (
    "encoding/json"
    "fmt"
    "io"
)

type VersionBuildResponse struct {
    Id      ProjectId `json:"project_id"`
    Name    string    `json:"project_name"`
    Version string    `json:"version"`
    Build
}

func (c *Config) GetVersionBuild(id ProjectId, version string, buildNumber int32) (VersionBuildResponse, error) {
    url := fmt.Sprintf("%s/projects/%s/versions/%s/builds/%d", c.BaseUrl, id, version, buildNumber)
    resp, err := c.Client.Get(url)
    if err != nil {
        return VersionBuildResponse{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return VersionBuildResponse{}, err
    }

    var v VersionBuildResponse
    err = json.Unmarshal(body, &v)
    if err != nil {
        return VersionBuildResponse{}, err
    }

    return v, nil
}
