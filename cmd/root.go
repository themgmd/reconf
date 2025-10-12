package cmd

import (
	"github.com/spf13/cobra"
	"github.com/themgmd/reconf/internal/constants"
	"os"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "reconf",
	Version: constants.CLIVersion,
	Short:   "A powerful CLI for app configuration",
	Long: `
A powerful CLI for app configuration.

Create configuration templates and variable-keys.
Store secrets in vault and present variables in golang types

A helpful documentation and next steps -> https://github.com/themgmd/reconf`,
}

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}
