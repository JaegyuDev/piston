package mojang_piston

import (
    "fmt"
    "github.com/cheggaaa/pb/v3"
    "io"
    "net/http"
    "os"
    "path"
)

type Download struct {
    doSnapshots bool
    url         string
    path        string
    version     string
}

func (d *Download) DoSnapshots(b bool) {
    d.doSnapshots = b
}

func (d *Download) Version(v string) {
    d.version = v

    if v == "latest" {
        versionManifest := GetPistonMeta()
        if len(versionManifest.Versions) == 0 {
            fmt.Println("No versions available.")
            os.Exit(1)
        }

        var latestVersion *Version
        for _, v := range versionManifest.Versions {
            if d.doSnapshots && v.Type == "snapshot" {
                latestVersion = &v
                break
            }

            if v.Type == "release" {
                latestVersion = &v
                break
            }
        }
        if latestVersion == nil {
            fmt.Println("No valid release versions found.")
            os.Exit(1)
        }
        d.version = latestVersion.ID
    }

    versionManifest := GetPistonMeta()
    var versionDataUrl string
    for _, v := range versionManifest.Versions {
        if v.ID == d.version {
            versionDataUrl = v.URL
            break
        }
    }

    if versionDataUrl == "" {
        fmt.Printf("Version '%s' not found.\n", d.version)
        return
    }

    pistonData := GetPistonData(versionDataUrl)
    serverJarURL := pistonData.Downloads.Server.URL
    if serverJarURL == "" {
        fmt.Printf("No server jar available for version '%s'.\n", d.version)
        return
    }

    d.url = serverJarURL
}

func (d *Download) Path(p string) {
    // todo: check if the target is a path or
    //  directory in case user wants to change name of path
    d.path = path.Join(p, "server.jar")
}

func (d *Download) Do() error {
    resp, err := http.Get(d.url)
    if err != nil {
        return fmt.Errorf("error while downloading %v ", err)
    }
    defer resp.Body.Close()

    file, err := os.Create(d.path)
    if err != nil {
        return fmt.Errorf("error while creating path %v", err)
    }

    defer file.Close()

    bar := pb.StartNew(int(resp.ContentLength))
    bar.SetMaxWidth(80)
    writer := bar.NewProxyWriter(file)

    _, err = io.Copy(writer, resp.Body)
    if err != nil {
        return fmt.Errorf("error while writing to disk %v", err)
    }

    bar.Finish()

    fmt.Printf("Downloaded server jar for version '%s' to '%s'.\n", d.version, d.path)

    return nil
}
