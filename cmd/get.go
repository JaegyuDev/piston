package cmd

import (
    "fmt"
    "github.com/JaegyuDev/piston-cli/internal/mojang-piston"
    "github.com/cheggaaa/pb/v3"
    "github.com/spf13/cobra"
    "io"
    "net/http"
    "os"
)

var getCmd = &cobra.Command{
    Use:   "get [version]",
    Short: "Download a Minecraft server jar for the specified version",
    Long: `Downloads the Minecraft server jar for the specified version from the Mojang Piston Meta API.
If a path is specified with the -o flag, the server jar will be saved there.
If no path is given, the jar will be downloaded to the current directory.`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        version := args[0]

        outputPath, err := cmd.Flags().GetString("output")
        if err != nil {
            panic(err)
        }

        versionManifest := mojang_piston.GetPistonMeta()
        var versionDataURL string
        for _, v := range versionManifest.Versions {
            if v.ID == version {
                versionDataURL = v.URL
                break
            }
        }

        if versionDataURL == "" {
            fmt.Printf("Version '%s' not found.\n", version)
            return
        }

        pistonData := mojang_piston.GetPistonData(versionDataURL)
        serverJarURL := pistonData.Downloads.Server.URL

        if serverJarURL == "" {
            fmt.Printf("No server jar available for version '%s'.\n", version)
            return
        }

        resp, err := http.Get(serverJarURL)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        jarFilePath := fmt.Sprintf("%s/server.jar", outputPath)
        file, err := os.Create(jarFilePath)
        if err != nil {
            panic(err)
        }
        defer file.Close()

        bar := pb.StartNew(int(resp.ContentLength))
        bar.SetMaxWidth(80)
        writer := bar.NewProxyWriter(file)

        _, err = io.Copy(writer, resp.Body)
        if err != nil {
            panic(err)
        }

        bar.Finish()

        fmt.Printf("Downloaded server jar for version '%s' to '%s'.\n", version, jarFilePath)
    },
}

func init() {
    getCmd.Flags().StringP("output", "o", ".", "Output path to save the server jar")
    rootCmd.AddCommand(getCmd)
}
