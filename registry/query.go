package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type regRespBody struct {
	Repos []string `json:"repositories"`
}

type repoRespBody struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func RegistryAPIGetAllRepos(endpoint string) ([]string, error) {
	url := RepoListURL(endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []string{}, fmt.Errorf("http response error code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var resp_body regRespBody
	if err = json.Unmarshal(body, &resp_body); err != nil {
		return []string{}, err
	}
	return resp_body.Repos, nil
}

func RegistryAPIGetAllTags(endpoint, repo string) ([]string, error) {
	url := TagListURL(endpoint, repo)
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []string{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var resp_body repoRespBody
	if err = json.Unmarshal(body, &resp_body); err != nil {
		return []string{}, err
	}
	return resp_body.Tags, nil
}

func ListRepo(endpoint string, repo string) ([]string, error) {
	image_full_names := make([]string, 0)
	tags, err := RegistryAPIGetAllTags(endpoint, repo)
	if err != nil {
		return image_full_names, err
	}
	repo_full_name := fmt.Sprintf("%s/%s", endpoint, repo)
	for _, tag := range tags {
		image_full_names = append(image_full_names, fmt.Sprintf("%s:%s", repo_full_name, tag))
	}
	return image_full_names, nil
}
