package v2

import "fmt"

// GetDownloadUrl returns the URL to download the server jar.
func (c *Config) GetDownloadUrl(id ProjectId, version string, build int32, download string) string {
    return fmt.Sprintf("%s/projects/%s/versions/%s/builds/%d/downloads/%s",
        ApiUrl,
        id,
        version,
        build,
        download,
    )
}
