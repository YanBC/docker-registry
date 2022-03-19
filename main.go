package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/YanBC/docker-query/registry"
)

func list(endpoint string) {
	names, err := registry.GetAllImageNames(endpoint)
	if err != nil {
		log.Fatalf("fail to list image repos, error: %s", err)
	}
	available_images := make([]string, 0)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, name := range names {
		wg.Add(1)
		go func(image_name string) {
			defer wg.Done()
			full_names, err := registry.GetImageWithTags(endpoint, image_name)
			if err != nil {
				log.Printf("fail   : %s", image_name)
			} else {
				mu.Lock()
				defer mu.Unlock()
				available_images = append(available_images, full_names...)
				log.Printf("succeed: %s", image_name)
			}
		}(name)
	}
	// print available images with tags
	wg.Wait()
	sort.Strings(available_images)
	fmt.Println(" ########################################### ")
	fmt.Println("Available images:")
	for _, full_name := range available_images {
		fmt.Println(full_name)
	}
}

func delete(image_repo string, tag string) {
	var msg string
	if tag != "" {
		image_full_name := fmt.Sprintf("%s:%s", image_repo, tag)
		if err := registry.DeleteImage(image_full_name); err != nil {
			msg = fmt.Sprintf("delete %s fail, %s", image_full_name, err)
		} else {
			msg = fmt.Sprintf("delete %s succeed", image_full_name)
		}
	} else {
		if err := registry.DeleteRepo(image_repo); err != nil {
			msg = fmt.Sprintf("delete %s fail, %s", image_repo, err)
		} else {
			msg = fmt.Sprintf("delete %s succeed", image_repo)
		}
	}
	fmt.Println(msg)
}

func help() {
	executable := os.Args[0]
	fmt.Printf("docker-query %s\n", VERSION)
	fmt.Printf("%s list:     list available images\n", executable)
	fmt.Printf("%s delete:   delete images\n", executable)
	fmt.Printf("Run %s list/delete -h for usage of each subcommand\n", executable)
}

func main() {
	// list
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	l_addr := listCmd.String("addr", "", "docker registry endpoint")

	// delete
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	d_repo := deleteCmd.String("repo", "", "image repo to delete")
	d_tag := deleteCmd.String("tag", "", "tag to delete, empty to delete all tags")

	// parse arguments
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		if *l_addr == "" {
			fmt.Println("invalid address")
			os.Exit(1)
		}
		list(*l_addr)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *d_repo == "" {
			fmt.Println("specify which image repo to delete with -repo")
			os.Exit(1)
		}
		delete(*d_repo, *d_tag)
	default:
		help()
		os.Exit(1)
	}
}
