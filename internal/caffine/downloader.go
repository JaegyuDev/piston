package caffine

import (
    "fmt"
    mojang_piston "github.com/JaegyuDev/piston/internal/mojang-piston"
    // papermc "github.com/JaegyuDev/piston/pkg/papermc/v2"
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
        return &mojang_piston.Download{}, nil
        // case "paper":
        // return &papermc.Download{}, nil
    default:
        return nil, fmt.Errorf("no loader found for version '%s'", name)
    }
}
