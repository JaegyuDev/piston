package cmd

import (
    "fmt"
    "github.com/JaegyuDev/piston/pkg/mojang-piston"
    "strings"

    "github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
    Use:   "search [query]",
    Short: "Search for a specific Minecraft version by name",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        allowSnapshots, err := cmd.Flags().GetBool("allow-snapshots")
        if err != nil {
            panic(err)
        }

        query := args[0]
        versionManifest := mojang_piston.GetPistonMeta()

        fmt.Printf("Search results for '%s':\n", query)
        found := false

        for _, v := range versionManifest.Versions {
            // currently no support for the pre 1.0 versions.
            if v.Type == "old_alpha" || v.Type == "old_beta" {
                continue
            }

            if !allowSnapshots && v.Type == "snapshot" {
                continue
            }

            if containsIgnoreCase(v.ID, query) {
                fmt.Printf(formatNormalizedOutput(v))
                found = true
            }
        }

        if !found {
            fmt.Println("No matching versions found.")
        }
    },
}

func containsIgnoreCase(str, substr string) bool {
    return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func init() {
    rootCmd.AddCommand(searchCmd)
}
