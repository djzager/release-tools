package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "log"

	"github.com/google/go-github/v49/github"
	"github.com/konveyor/release-tools/action"
)

func main() {

	ghContext, err := action.GetGitHubContext()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the event
	event, err := func() (github.PullRequestEvent, error) {
		eventFile, err := os.Open(ghContext.GithubEventPath)
		if err != nil {
			return github.PullRequestEvent{}, fmt.Errorf("unable to load event file: %w", err)
		}
		defer func() {
			// As we are not writing to the file, we can omit the error
			_ = eventFile.Close()
		}()

		var event github.PullRequestEvent
		if err := json.NewDecoder(eventFile).Decode(&event); err != nil {
			return event, fmt.Errorf("unable to unmarshal event: %w", err)
		}
		return event, nil
	}()

	fmt.Printf("PR Title %v", event.PullRequest.Title)
}
