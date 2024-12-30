package internal

type Resource struct {
	GID          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type Data []*Resource

type Response struct {
	Data Data `json:"data"`
}
