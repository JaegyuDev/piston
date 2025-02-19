package cmd

import (
    "fmt"
    "github.com/JaegyuDev/piston/internal/caffine"
    "github.com/spf13/cobra"
    "os"
    "path"
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

        loader, err := cmd.Flags().GetString("loader")
        if err != nil {
            fmt.Println("Error getting loader parameter:", err)
            os.Exit(1)
        }

        downloader, err := caffine.GetLoader(loader)
        if err != nil {
            fmt.Println("Error getting loader: ", err)
            os.Exit(1)
        }
        downloader.Version(version)

        outputPath, err := cmd.Flags().GetString("output")
        if err != nil {
            fmt.Println("Error getting output path: ", err)
            os.Exit(1)
        }
        downloader.Path(path.Join(outputPath, "server.jar"))

        doSnapshots, err := cmd.Flags().GetBool("allow-snapshots")
        if err != nil {
            fmt.Println("Error getting allow-snapshots: ", err)
            os.Exit(1)
        }
        downloader.DoSnapshots(doSnapshots)

        err = downloader.Do()
        if err != nil {
            fmt.Println("Error downloading: ", err)
            os.Exit(1)
        }
    },
}

func init() {
    getCmd.Flags().StringP("output", "o", ".", "Output path to save the server jar")
    getCmd.Flags().StringP("loader", "l", "vanilla", "Set the desired server software")
    rootCmd.AddCommand(getCmd)
}
