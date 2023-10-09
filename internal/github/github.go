package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GitHubUser struct {
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Url    string `json:"html_url"`
	Avatar string `json:"avatar_url"`
}

func User(username string) GitHubUser {
	// set up the GitHub API endpoint
	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	log.Println("Loading profile", url)

	// make a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to retrieve profile data: %d", resp.StatusCode)
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Failed to parse body: %v", err)
	}

	// Parse the JSON response
	var user GitHubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Failed to unmarshal body: %v", err)
	}

	return user
}
