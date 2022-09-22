/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Will create a new [subcommand]. (for example a job)",
}

func init() {
	rootCmd.AddCommand(newCmd)
}
