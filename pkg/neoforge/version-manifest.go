package neoforge

import (
    "encoding/xml"
    "net/http"
)

const BaseURL = "https://maven.neoforged.net/releases/net/neoforged/neoforge/"

type Manifest struct {
    Metadata struct {
        GroupId    string `xml:"groupId"`
        ArtifactId string `xml:"artifactId"`
        Versioning struct {
            Latest      string   `xml:"latest"`
            Release     string   `xml:"release"`
            Versions    []string `xml:"versions"`
            LastUpdated string   `xml:"lastUpdated"`
        } `xml:"versioning"`
    } `xml:"metadata"`
}

func GetManifest() (*Manifest, error) {
    resp, err := http.Get(BaseURL + "maven-metadata.xml")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var manifest Manifest
    err = xml.NewDecoder(resp.Body).Decode(&manifest)
    if err != nil {
        return nil, err
    }

    return &manifest, nil
}
