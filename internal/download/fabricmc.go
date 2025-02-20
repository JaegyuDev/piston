package download

import (
    "fmt"
    fabricmc "github.com/JaegyuDev/piston/pkg/fabricmc/v2"
    "os"
)

type FabricMC struct {
    doSnapshots bool
    path        string
    gameVersion string
}

func (f *FabricMC) DoSnapshots(b bool) {
    f.doSnapshots = b
}

func (f *FabricMC) Path(path string) {
    f.path = path
}

func (f *FabricMC) Version(version string) {
    client := fabricmc.NewConfig()

    if version == "latest" {
        gameVersion, err := client.GetLatestGameVersion(f.doSnapshots)
        if err != nil {
            fmt.Println("Failed to get latest game version for FabricMC: ", err)
            os.Exit(1)
        }

        f.gameVersion = gameVersion
        return
    }

    if ok, _ := client.GetGameVersion(version); !ok {
        fmt.Printf("Version %s is not supported by Fabric yet!\n", version)
        os.Exit(1)
    }

    f.gameVersion = version
}

func (f *FabricMC) Do() error {
    client := fabricmc.NewConfig()

    installerVer, err := client.GetLatestInstallerVersion()
    if err != nil {
        return err
    }

    loaderVer, err := client.GetLatestLoaderVersion()
    if err != nil {
        return err
    }

    url := client.GetDownloadUrl(f.gameVersion, loaderVer, installerVer)
    err = Download(url, f.path)
    if err != nil {
        return err
    }

    fmt.Printf("Downloaded fabric server jar for game version '%s', loader version '%s', and installer version '%s'"+
        "to '%s'.\n", f.gameVersion, loaderVer, installerVer, f.path)

    return nil
}
