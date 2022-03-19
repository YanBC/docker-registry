package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repo struct {
	Images []string `json:"repositories"`
}

type Image struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func QueryRepo(endpoint string) (Repo, error) {
	url := fmt.Sprintf("http://%s/v2/_catalog", endpoint)
	resp, err := http.Get(url)
	if err != nil {
		return Repo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Repo{}, fmt.Errorf("http response error code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Repo{}, err
	}
	var repo Repo
	if err = json.Unmarshal(body, &repo); err != nil {
		return Repo{}, err
	}
	return repo, nil
}

func QueryImage(endpoint string, name string) (Image, error) {
	url := fmt.Sprintf("http://%s/v2/%s/tags/list", endpoint, name)
	resp, err := http.Get(url)
	if err != nil {
		return Image{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return Image{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Image{}, err
	}
	var image Image
	if err = json.Unmarshal(body, &image); err != nil {
		return Image{}, err
	}
	return image, nil
}

func GetAllImageNames(endpoint string) ([]string, error) {
	repo, err := QueryRepo(endpoint)
	if err != nil {
		return []string{}, err
	}
	return repo.Images, nil
}

func GetImageWithTags(endpoint string, image_name string) ([]string, error) {
	full_name_with_tags := make([]string, 0)
	image, err := QueryImage(endpoint, image_name)
	if err != nil {
		return full_name_with_tags, err
	}
	full_name := fmt.Sprintf("%s/%s", endpoint, image_name)
	for _, tag := range image.Tags {
		full_name_with_tags = append(full_name_with_tags, fmt.Sprintf("%s:%s", full_name, tag))
	}
	return full_name_with_tags, nil
}
