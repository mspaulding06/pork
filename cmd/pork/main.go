package main

import (
	"os"

	"github.com/mspaulding06/pork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	rootCmd.AddCommand(pork.ForkCmd)
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("pork")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
