package download

import (
    "fmt"
    papermc "github.com/JaegyuDev/piston/pkg/papermc/v2"
    "os"
)

type PaperMC struct {
    project     papermc.ProjectId
    doSnapshots bool
    path        string
    version     string
}

func (p *PaperMC) Project(l string) {
    p.project = papermc.ProjectId(l)
}

func (p *PaperMC) DoSnapshots(b bool) {
    p.doSnapshots = b
}

func (p *PaperMC) Version(v string) {
    if ok, _ := GetVersion(v); !ok {
        fmt.Printf("Version %s not found\n", v)
        os.Exit(1)
    }

    p.version = v
}

func (p *PaperMC) Path(path string) {
    p.path = path
}

func (p *PaperMC) Do() error {
    // first we need to retrieve a list of builds, which also allows us to get the downloads too.
    client := papermc.NewConfig()

    versionResp, err := client.GetVersionBuilds(p.project, p.version)
    if err != nil {
        fmt.Printf("GetVersionBuilds error: %s\n", err)
        os.Exit(1)
    }

    build := versionResp.Builds[len(versionResp.Builds)-1]
    buildName := build.Downloads.Application.Name

    url := client.GetDownloadUrl(p.project, p.version, build.Build, buildName)

    err = Download(url, p.path)
    if err != nil {
        return err
    }

    fmt.Printf("Downloaded %s server jar for version '%s' to '%s'.\n", versionResp.Name, p.version, p.path)

    return nil
}
