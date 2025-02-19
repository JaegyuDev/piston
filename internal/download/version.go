package download

import (
    mojang_piston "github.com/JaegyuDev/piston/pkg/mojang-piston"
)

func GetVersion(version string) (bool, string) {
    versionManifest := mojang_piston.GetPistonMeta()
    var versionDataUrl string
    for _, v := range versionManifest.Versions {
        if v.ID == version {
            versionDataUrl = v.URL
            break
        }
    }

    if versionDataUrl == "" {
        return false, versionDataUrl
    }

    return true, versionDataUrl
}

func GetLatestVersion(snapshot bool) string {
    versionManifest := mojang_piston.GetPistonMeta()
    if snapshot {
        return versionManifest.Latest.Snapshot
    }

    return versionManifest.Latest.Release
}
