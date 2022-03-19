# Docker Registry Querier
A handy tool to query your self-hosted docker registry.
Subcommand:
- `list`: list all available images
- `delete`: delete specified image

Build: `go build -o docker-registry main.go version.go`

Usage:
```bash
# print usage
./docker-registry help
# output:
# docker-query 1.0.0
# ./docker-registry list:     list available images
# ./docker-registry delete:   delete images
# Run ./docker-registry list/delete -h for usage of each subcommand

# list registry
./docker-registry list -addr 172.17.0.2:5000

# delete image(s)
## delete the whole image repo
./docker-registry delete -repo 172.17.0.2:5000/busybox
## delete single image
./docker-registry delete -repo 172.17.0.2:5000/busybox -tag latest
```


## TODOs
1. Regex in `seperate_image_full_name` and `seperate_repo_full_name`
2. Add logging level option
3. <del>Refactor, rename functions and variables</del>
4. Fix bug in delete. When deleting images with the same digest, latter delete will fail with 404 http code.


## Misc
### Naming convention
Normally, a self-hosted docker **registry** contains many image repositories (**repo**). And each image repository contains one image with multiple tags. One image repo name and one tag identify one available **image**.

In that regard, when naming variables in this project, I use the following conventions:
1. `registry`/`endpoint`: `<ip>:<port>`, address of a docker registry
2. `repo`: name of a repo
3. `repo_full_name`: `<endpoint>/<repo>`
4. `tag`: one of possibly many tags of a repo
5. `image`: `<repo>:<tag>`
6. `image_full_name`: `<endpoint>/<repo>:<tag>`
