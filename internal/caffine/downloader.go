package caffine

import (
    "fmt"
    "github.com/JaegyuDev/piston/internal/download"
    "strings"
)

type Downloader interface {
    DoSnapshots(bool)
    Version(string)
    Path(string)
    Do() error
}

func GetLoader(name string) (Downloader, error) {
    switch strings.ToLower(name) {
    case "vanilla", "piston", "mojang":
        return &download.Vanilla{}, nil
    case "paper", "travertine", "waterfall", "velocity", "folia":
        downloader := download.PaperMC{}
        downloader.Project(name)

        return &downloader, nil
    case "fabric":
        return &download.FabricMC{}, nil

    default:
        return nil, fmt.Errorf("no loader found for version '%s'", name)
    }
}
