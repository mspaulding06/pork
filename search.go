package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryList := SearchByKeyword(args)
		for _, repository := range repositoryList {
			fmt.Println(repository)
		}
	},
}

func SearchByKeyword(keywords []string) []string {
	return []string{"myrepository"}
}
