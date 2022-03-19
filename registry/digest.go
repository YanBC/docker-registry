package registry

import (
	"fmt"
	"net/http"
)

func RegistryAPIGetImageDigest(endpoint string, repo string, tag string) (string, error) {
	// Create request
	url := ImageDigestURL(endpoint, repo, tag)
	req, _ := http.NewRequest("HEAD", url, nil)
	// See https://docs.docker.com/registry/spec/api/#deleting-an-image
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	client.CloseIdleConnections()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("http response error code: %d", resp.StatusCode)
	}

	// Read http header
	digest := resp.Header.Get("Docker-Content-Digest")
	if digest == "" {
		return digest, fmt.Errorf("registry return empty digest, url: %s", url)
	}

	return digest, nil
}
