package download

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type NeoForge struct {
    doSnapshots bool
    version     string
    path        string
}

func (n *NeoForge) DoSnapshots(b bool) {
    n.doSnapshots = b
}

func (n *NeoForge) Version(s string) {
    triplet := SplitTriplet(s)

    if triplet[1] < 20 || triplet[1] == 20 && triplet[2] < 2 {
        fmt.Printf("Neoforge only supports versions 1.20.2 and up.")
        os.Exit(1)
    }

    // should be okay to store the version here now
    n.version = s
}

func (n *NeoForge) Path(s string) {
    // TODO implement me
    panic("implement me")
}

func (n *NeoForge) Do() error {
    // TODO implement me
    panic("implement me")
}

func SplitTriplet(s string) []int {
    subs := strings.SplitN(s, ".", 3)
    parts := make([]int, len(subs))

    var err error
    parts[0], err = strconv.Atoi(subs[0])
    if err != nil {
        panic(err)
    }

    parts[1], err = strconv.Atoi(subs[1])
    if err != nil {
        panic(err)
    }

    // makes sure to drop any neoforge release type data from triplet
    subs2 := strings.Split(subs[2], "-")

    parts[2], err = strconv.Atoi(subs2[0])
    if err != nil {
        panic(err)
    }

    return parts
}
