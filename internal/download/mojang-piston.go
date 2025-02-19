// TODO: probably should be handling the errors further down the stack, probably won't fix this right away tho since
//  it doesn't really change much of the functionality

package download

import (
    "fmt"
    mojang_piston "github.com/JaegyuDev/piston/pkg/mojang-piston"
)

type Vanilla struct {
    doSnapshots bool
    path        string
    version     string
}

func (d *Vanilla) DoSnapshots(b bool) {
    d.doSnapshots = b
}

func (d *Vanilla) Version(v string) {
    d.version = v

    if v == "latest" {
        d.version = GetLatestVersion(d.doSnapshots)
    }
}

func (d *Vanilla) Path(p string) {
    // todo: check if the target is a path or
    //  directory in case user wants to change name of path
    d.path = p
}

func (d *Vanilla) Do() error {
    var (
        ok             bool
        versionDataUrl string
    )

    if ok, versionDataUrl = GetVersion(d.version); !ok {
        return fmt.Errorf("Version %s is not available\n", d.version)
    }

    pistonData := mojang_piston.GetPistonData(versionDataUrl)
    serverJarURL := pistonData.Downloads.Server.URL
    if serverJarURL == "" {
        return fmt.Errorf("No server jar available for version '%s'.\n", d.version)
    }

    err := Download(serverJarURL, d.path)
    if err != nil {
        return err
    }

    fmt.Printf("Downloaded vanilla server jar for version '%s' to '%s'.\n", d.version, d.path)

    return nil
}
