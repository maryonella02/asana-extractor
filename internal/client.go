package internal

import (
	config "asana-extractor/cmd/config"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const Users = "users"
const Projects = "projects"

type AsanaClient struct {
	token      string
	httpClient *http.Client
	//rate limiter
}

func NewClient(cfg *config.Config) *AsanaClient {
	return &AsanaClient{
		token: cfg.PAT,
		httpClient: &http.Client{
			Timeout: time.Duration(1) * time.Second,
		},
	}
}

func (a *AsanaClient) GetResource(path string) ([]byte, error) {
	// TODO rate limit

	//URL should be extracted to config file or passed to function
	req, err := http.NewRequest("GET", fmt.Sprintf("https://app.asana.com/api/1.0/%s", path), nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token))

	resp, err := a.httpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)

	return body, nil
}

func (a *AsanaClient) GetUsers() (Data, error) {
	bytes, err := a.GetResource(Users)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s : %w", Users, err)
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Data, nil

}
func (a *AsanaClient) GetProjects() (Data, error) {
	bytes, err := a.GetResource(Projects)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s : %w", Projects, err)
	}

	var response Response
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Data, nil

}
func (a *AsanaClient) ExportUsersToFile() error {
	data, err := a.GetUsers()
	if err != nil {
		return fmt.Errorf("error export: %w", err)
	}
	err = writeToFile(Users, data)
	if err != nil {
		return fmt.Errorf("error export: %w", err)
	}

	return nil

}
func (a *AsanaClient) ExportProjectsToFile() error {
	data, err := a.GetProjects()
	if err != nil {
		return fmt.Errorf("error export: %w", err)
	}
	err = writeToFile(Projects, data)
	if err != nil {
		return fmt.Errorf("error export: %w", err)
	}

	return nil

}
func writeToFile(path string, data Data) error {

	for _, resource := range data {
		content, err := json.Marshal(&resource)
		if err != nil {
			return fmt.Errorf("failed to marshal %s: %w", path, err)
		}

		name := fmt.Sprintf("/output/%s/%s", path, resource.GID)
		if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {

			err = os.WriteFile(name, content, 0664)
			if err != nil {
				return fmt.Errorf("failed to write to file %s: %w", path, err)

			}

		}
	}
	return nil

}
