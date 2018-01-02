package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ReleasesInfo Object for release info
type ReleasesInfo struct {
	ID      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

// ReleaseInfoer TODO
type ReleaseInfoer interface {
	GetLatestReleaseTag(string) (string, error)
}

// GithubReleaseInfoer struct
type GithubReleaseInfoer struct{}

// GetLatestReleaseTag func get body response and print JSON
func (gh GithubReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	releases := []ReleasesInfo{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	fmt.Println(releases)

	tag := releases[0].TagName

	return tag, nil
}

// Function to get the message to display to the end user.
func getReleaseTagMessage(ri ReleaseInfoer, repo string) (string, error) {
	tag, err := ri.GetLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying GitHub API: %s", err)
	}

	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	gh := GithubReleaseInfoer{}
	msg, err := getReleaseTagMessage(gh, "lindep/golang_exp")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(msg)
}
