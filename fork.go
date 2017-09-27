package pork

import (
	"log"

	"github.com/spf13/cobra"
)

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("You must supply a repository")
		}
		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("Unable to fork repository:", err)
		}
	},
}

func ForkRepository(repository string) error {
	return nil
}
