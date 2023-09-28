# Docker Registry Querier
A handy tool to query your self-hosted docker registry.
Subcommand:
- `list`: list all available images
- `delete`: delete specified image

## Install
```bash
git clone https://github.com/YanBC/docker-registry.git
cd docker-registry
go install github.com/yanbc/docker-registry
```

## Usage

To view the available commands and options, run:

```bash
docker-registry help
```

### List images

To list all the images available in a registry, run:

```bash
docker-registry list -addr <registry_address>
```

Replace `<registry_address>` with the IP address and port number of your registry.

### Delete images

To delete an image, run:

```bash
docker-registry delete -repo <repo_full_name> [-tag <tag>]
```

Replace `<repo_full_name>` with the full name of the repository you want to delete, in the format `<registry_address>/<repo_name>`. If you want to delete a specific tag of an image, add the `-tag` option.

## TODOs
1. Regex in `seperate_image_full_name` and `seperate_repo_full_name`
2. Add logging level option
3. <del>Refactor, rename functions and variables</del>
4. Fix bug in delete. When deleting images with the same digest, latter delete will fail with 404 http code.
5. Add github actions for auto build and release.


## Naming convention

In a self-hosted Docker registry, there are many image repositories (**repo**), each containing one or more images with multiple tags. A repository name and a tag together identify a specific **image**.

In this tool, we use the following naming conventions for variables:

1. `registry` or `endpoint`: The address (IP and port) of a Docker registry.
2. `repo`: The name of a repository.
3. `repo_full_name`: The full name of a repository, in the format `<registry_address>/<repo_name>`.
4. `tag`: A specific tag associated with an image.
5. `image`: The name of an image, in the format `<repo>:<tag>`.
6. `image_full_name`: The full name of an image, in the format `<registry_address>/<repo>:<tag>`.
