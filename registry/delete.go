package registry

import (
	"fmt"
	"net/http"
)

func seperate_image_full_name(image_full_name string) (string, string, string) {
	endpoint, repo, tag := "", "", ""
	sep1, sep2 := 0, 0
	for i := 0; i < len(image_full_name); i++ {
		if image_full_name[i] == '/' {
			sep1 = i
			break
		}
	}
	for j := sep1; j < len(image_full_name); j++ {
		if image_full_name[j] == ':' {
			sep2 = j
			break
		}
	}

	if sep1 != 0 && sep2 != 0 {
		endpoint = image_full_name[:sep1]
		repo = image_full_name[sep1+1 : sep2]
		tag = image_full_name[sep2+1:]
	}
	return endpoint, repo, tag
}

func seperate_repo_full_name(repo_full_name string) (string, string) {
	endpoint, repo := "", ""
	var sep int
	for sep = 0; sep < len(repo_full_name); sep++ {
		if repo_full_name[sep] == '/' {
			break
		}
	}

	if sep != 0 {
		endpoint = repo_full_name[:sep]
		repo = repo_full_name[sep+1:]
	}
	return endpoint, repo
}

func RegistryAPIDeleteImage(endpoint string, repo string, tag string) error {
	digest, err := RegistryAPIGetImageDigest(endpoint, repo, tag)
	if err != nil {
		return err
	}

	// Create request
	url := ImageDeleteURL(endpoint, repo, digest)
	req, _ := http.NewRequest("DELETE", url, nil)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	client.CloseIdleConnections()
	if resp.StatusCode != 202 {
		return fmt.Errorf("http response error code: %d", resp.StatusCode)
	}

	return nil
}

func DeleteImage(image_full_name string) error {
	endpoint, repo, tag := seperate_image_full_name(image_full_name)
	if tag == "" {
		return fmt.Errorf("invalid image name: %s", image_full_name)
	}

	return RegistryAPIDeleteImage(endpoint, repo, tag)
}

func DeleteRepo(repo_full_name string) error {
	endpoint, repo := seperate_repo_full_name(repo_full_name)
	tags, err := RegistryAPIGetAllTags(endpoint, repo)
	if err != nil {
		return err
	}

	all_succeed := true
	var ret_err error
	for _, tag := range tags {
		if err := RegistryAPIDeleteImage(endpoint, repo, tag); err != nil {
			all_succeed = false
			ret_err = err
		}
	}
	if !all_succeed {
		return fmt.Errorf("not all delete succeeds, error: %s", ret_err)
	}

	return nil
}
