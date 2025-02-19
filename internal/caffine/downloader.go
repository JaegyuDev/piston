package caffine

import (
    "fmt"
    "github.com/JaegyuDev/piston/internal/download"
)

type Downloader interface {
    DoSnapshots(bool)
    Version(string)
    Path(string)
    Do() error
}

func GetLoader(name string) (Downloader, error) {
    switch name {
    case "vanilla", "piston", "mojang":
        return &download.Vanilla{}, nil
    case "paper", "travertine", "waterfall", "velocity", "folia":
        downloader := download.PaperMC{}
        downloader.Project(name)

        return &downloader, nil

    default:
        return nil, fmt.Errorf("no loader found for version '%s'", name)
    }
}
