package internal

import (
	"asana-extractor/cmd"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AsanaClient struct {
	token      string
	httpClient *http.Client
	//rate limiter
}

func NewClient(cfg *cmd.Config) *AsanaClient {
	return &AsanaClient{
		token: cfg.PAT,
		httpClient: &http.Client{
			Timeout: time.Duration(1) * time.Second,
		},
	}
}

func (a *AsanaClient) GetUsers() {
	// TODO rate limit

	//URL should be extracted to config file or passed to function
	req, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/users", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)
}

func (a *AsanaClient) GetProjects() {
	// TODO rate limit

	// TODO rate limit

	//URL should be extracted to config file or passed to function
	req, err := http.NewRequest("GET", "https://app.asana.com/api/1.0/projects", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

}
