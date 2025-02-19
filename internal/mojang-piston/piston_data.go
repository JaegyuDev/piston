package mojang_piston

import (
    "encoding/json"
    "io"
    "net/http"
    "time"
)

type PistonData struct {
    ID                     string    `json:"id"`
    Type                   string    `json:"type"`
    MainClass              string    `json:"mainClass"`
    ReleaseTime            time.Time `json:"releaseTime"`
    Assets                 string    `json:"assets"`
    ComplianceLevel        int       `json:"complianceLevel"`
    MinimumLauncherVersion int       `json:"minimumLauncherVersion"`

    AssetIndex struct {
        ID  string `json:"id"`
        URL string `json:"url"`
    } `json:"assetIndex"`

    Downloads struct {
        Client struct {
            URL string `json:"url"`
        } `json:"client"`
        Server struct {
            URL string `json:"url"`
        } `json:"server"`
    } `json:"downloads"`

    Libraries []struct {
        Name string `json:"name"`
    } `json:"libraries"`

    Logging struct {
        Client struct {
            Argument string `json:"argument"`
        } `json:"client"`
    } `json:"logging"`
}

func GetPistonData(url string) PistonData {
    resp, err := http.Get(url)
    if err != nil {
        panic(err) // this shouldn't panic unless there isn't a connection.
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var pistonData PistonData
    err = json.Unmarshal(body, &pistonData)
    if err != nil {
        panic(err)
    }

    return pistonData
}
