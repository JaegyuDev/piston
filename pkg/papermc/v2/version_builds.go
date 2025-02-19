package v2

import (
    "encoding/json"
    "fmt"
    "io"
    "time"
)

type Application struct {
    Name   string `json:"name"`
    Sha256 string `json:"sha256"`
}

type Change struct {
    Commit  string `json:"commit"`
    Summary string `json:"summary"`
    Message string `json:"message"`
}

type Build struct {
    Build int32     `json:"build"`
    Time  time.Time `json:"time"`
    // Channel either {default, experimental}
    Channel   string   `json:"channel"`
    Promoted  bool     `json:"promoted"`
    Changes   []Change `json:"changes"`
    Downloads struct {
        Application Application `json:"application"`
    } `json:"downloads"`
}

type VersionBuildsResponse struct {
    Id      ProjectId `json:"project_id"`
    Name    string    `json:"project_name"`
    Version string    `json:"version"`
    Builds  []Build   `json:"builds"`
}

func (c *Config) GetVersionBuilds(id ProjectId, version string) (VersionBuildsResponse, error) {
    url := fmt.Sprintf("%s/projects/%s/versions/%s/builds", c.BaseUrl, id, version)
    resp, err := c.Client.Get(url)
    if err != nil {
        return VersionBuildsResponse{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return VersionBuildsResponse{}, err
    }

    var v VersionBuildsResponse
    err = json.Unmarshal(body, &v)
    if err != nil {
        return VersionBuildsResponse{}, err
    }

    return v, nil
}
