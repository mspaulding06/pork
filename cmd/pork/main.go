package main

import (
	"github.com/mspaulding06/pork"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "Project Forking Tool for GitHub",
	}
	rootCmd.AddCommand(pork.SearchCmd)
	rootCmd.AddCommand(pork.DocsCmd)
	rootCmd.AddCommand(pork.CloneCmd)
}
