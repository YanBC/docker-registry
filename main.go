package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"sync"
)

type Repo struct {
	Images []string `json:"repositories"`
}

type Image struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

var available_images []string
var mu sync.Mutex
var wg sync.WaitGroup

func query_repo(endpoint string) (Repo, error) {
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

func query_image(endpoint string, name string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/v2/%s/tags/list", endpoint, name))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("http response error code: %d", resp.StatusCode))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var image Image
	if err = json.Unmarshal(body, &image); err != nil {
		panic(err)
	}

	image_name := fmt.Sprintf("%s/%s", endpoint, name)
	mu.Lock()
	defer mu.Unlock()
	for _, tag := range image.Tags {
		available_images = append(available_images, fmt.Sprintf("%s:%s", image_name, tag))
	}
	wg.Done()
}

func main() {
	addr := flag.String("addr", "", "docker registry endpoint")
	flag.Parse()

	if *addr == "" {
		fmt.Println("invalid address")
		os.Exit(64)
	}

	repo, err := query_repo(*addr)
	if err != nil {
		fmt.Println("fail to query available images")
		panic(err)
	}

	available_images = make([]string, 0)
	for _, image_name := range repo.Images {
		wg.Add(1)
		query_image(*addr, image_name)
	}

	wg.Wait()
	sort.Strings(available_images)
	for _, full_name := range available_images {
		fmt.Println(full_name)
	}
}