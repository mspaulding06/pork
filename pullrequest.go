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

var (
	destRepo           string
	sourceRepo         string
	pullRequestTitle   string
	pullRequestMessage string
)

type PullRequestPayload struct {
	Title        string `json:"title"`
	Message      string `json:"body"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       bool   `json:"maintainer_can_modify"`
}

type PullRequestResponse struct {
	URL string `json:"html_url"`
}

var PullRequestCmd = &cobra.Command{
	Use:   "pullrequest",
	Short: "Create a Pull Request",
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreatePullRequest(); err != nil {
			log.Fatalln("Failed to create pull request:", err)
		}
	},
}

func CreatePullRequest() error {
	sourceValues := strings.Split(sourceRepo, ":")
	if !(len(sourceValues) == 1 || len(sourceValues) == 2) {
		return fmt.Errorf("Source repository must in the format [owner:]branch got %v", sourceRepo)
	}
	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		return fmt.Errorf("Destination repository must be in the format owner/project:branch got %v", destRepo)
	}
	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		return fmt.Errorf("Destination repository must be in the format owner/project:branch got %v", destRepo)
	}
	payload := &PullRequestPayload{
		Title:        pullRequestTitle,
		Message:      pullRequestMessage,
		SourceBranch: sourceRepo,
		DestBranch:   destBranchValues[1],
		Modify:       true,
	}
	return GitHubAPI().Call("pullrequest", map[string]string{
		"owner":   destValues[0],
		"project": destValues[1],
	}, payload)
}

func PullRequestSuccess(resp *http.Response) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	respContent := PullRequestResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("Created Pull Request: %s\n", respContent.URL)
	return nil
}

func PullRequestDefaultRouter(resp *http.Response) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetPullRequestResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(201, PullRequestSuccess)
	router.DefaultRouter = PullRequestDefaultRouter
	resource := nap.NewResource("/repos/{{.owner}}/{{.project}}/pulls", "POST", router)
	return resource
}

func init() {
	PullRequestCmd.Flags().StringVarP(&sourceRepo, "source", "s", "", "source repository")
	PullRequestCmd.Flags().StringVarP(&destRepo, "destination", "d", "", "destination repository")
	PullRequestCmd.Flags().StringVarP(&pullRequestTitle, "title", "t", "Basic Pull Request", "pull request title")
	PullRequestCmd.Flags().StringVarP(&pullRequestMessage, "message", "m", "", "pull request message")
}
