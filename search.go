package pork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mspaulding06/nap"
	"github.com/spf13/cobra"
)

type SearchResponse struct {
	Results []*SearchResult `json:"items"`
}

type SearchResult struct {
	FullName string `json:"full_name"`
}

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for GitHub repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		if err := SearchByKeyword(args); err != nil {
			log.Fatalln("Search Failed:", err)
		}
	},
}

func SearchByKeyword(keywords []string) error {
	return GitHubAPI().Call("search", map[string]string{
		"query": strings.Join(keywords, "+"),
	})
}

func SearchSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := SearchResponse{}
	json.Unmarshal(content, &respContent)
	for _, item := range respContent.Results {
		fmt.Println(item.FullName)
	}
	return nil
}

func SearchDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetSearchResource() *nap.RestResource {
	searchRouter := nap.NewRouter()
	searchRouter.DefaultRouter = SearchDefaultRouter
	searchRouter.RegisterFunc(200, SearchSuccess)
	search := nap.NewResource("/search/repositories?q={{.query}}", "GET", searchRouter)
	return search
}
