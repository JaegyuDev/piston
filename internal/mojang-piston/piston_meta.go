package mojang_piston

import (
    "encoding/json"
    "io"
    "net/http"
    "time"
)

const manifestURL = "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json"

type PistonMeta struct {
    Latest struct {
        Release  string `json:"release"`
        Snapshot string `json:"snapshot"`
    } `json:"latest"`
    Versions []Version `json:"versions"`
}

type Version struct {
    ID              string    `json:"id"`
    Type            string    `json:"type"`
    URL             string    `json:"url"`
    Time            time.Time `json:"time"`
    ReleaseTime     string    `json:"releaseTime"`
    Sha1            string    `json:"sha1"`
    ComplianceLevel int       `json:"complianceLevel"`
}

func GetPistonMeta() PistonMeta {
    resp, err := http.Get(manifestURL)
    if err != nil {
        panic(err) // this shouldnt panic unless there isnt a connection.
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var pistonMeta PistonMeta
    err = json.Unmarshal(body, &pistonMeta)
    if err != nil {
        panic(err)
    }

    return pistonMeta
}
