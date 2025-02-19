package download

import (
    "fmt"
    "github.com/cheggaaa/pb/v3"
    "io"
    "net/http"
    "os"
)

func Download(url, path string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("error while downloading %v ", err)
    }
    defer resp.Body.Close()

    file, err := os.Create(path)
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

    return nil
}
