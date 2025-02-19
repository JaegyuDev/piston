package cmd

import (
    "fmt"
    "github.com/JaegyuDev/piston/pkg/mojang-piston"
    "os"

    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "piston",
    Short: "A CLI tool for interacting with Mojang's Piston API",
    Long:  `piston is a command-line utility for downloading Minecraft server versions via Mojang's Piston API.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func formatNormalizedOutput(v mojang_piston.Version) string {
    return fmt.Sprintf("\t%-*s %s\n", 10, fmt.Sprintf("(%s)", v.Type), v.ID)
}

func init() {
    rootCmd.PersistentFlags().BoolP("allow-snapshots", "s", false, "Allow snapshots to be returned from the APIs. Doesn't affect __direct__ calls, like get [snapshot].")
}
