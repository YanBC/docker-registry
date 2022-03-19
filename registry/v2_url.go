package registry

import "fmt"

func TagListURL(endpoint string, image_repo string) string {
	return fmt.Sprintf("http://%s/v2/%s/tags/list", endpoint, image_repo)
}

func RepoListURL(endpoint string) string {
	return fmt.Sprintf("http://%s/v2/_catalog", endpoint)
}

func ImageDigestURL(endpoint string, repo string, tag string) string {
	return fmt.Sprintf("http://%s/v2/%s/manifests/%s", endpoint, repo, tag)
}

func ImageDeleteURL(endpoint string, repo string, digest string) string {
	return fmt.Sprintf("http://%s/v2/%s/manifests/%s", endpoint, repo, digest)
}
