package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		fmt.Printf("get image list failed, err: %v", err)
		os.Exit(2)
	}
	// fmt.Printf("ContainerList: %v\n", containers)
	for _, image := range images {
		fmt.Println(image.ID, image.RepoTags)
	}
	fmt.Println()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	// fmt.Printf("ContainerList: %v\n", containers)
	for _, container := range containers {
		fmt.Println(container.ID, container.Names, container.Status)
	}
}
