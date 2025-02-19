package cmd

import (
    "fmt"
    "github.com/JaegyuDev/piston/pkg/mojang-piston"

    "github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all available Minecraft versions",
    Long:  "Lists all available Minecraft versions retrieved from the Mojang Piston Meta API.",
    Run: func(cmd *cobra.Command, args []string) {
        allowSnapshots, err := cmd.Flags().GetBool("allow-snapshots")
        if err != nil {
            panic(err)
        }

        versionManifest := mojang_piston.GetPistonMeta()

        fmt.Println("Available versions:")
        for _, v := range versionManifest.Versions {
            // currently no support for the pre 1.0 versions.
            if v.Type == "old_alpha" || v.Type == "old_beta" {
                continue
            }

            if !allowSnapshots && v.Type == "snapshot" {
                continue
            }

            fmt.Printf(formatNormalizedOutput(v))
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
