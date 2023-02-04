package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os/exec"
)

func IsDockerRunning() bool {
	cmdStruct := exec.Command("docker", "info")
	_, err := cmdStruct.Output()
	return err == nil
}

func GetContainers() []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	return containers
}

func DeleteContainer(container types.Container) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	fmt.Printf("Stopping container %s...\n", container.Names)
	if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
		panic(err)
	}
	fmt.Printf("Deleting container %s...\n", container.Names)
	if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}); err != nil {
		panic(err)
	}
}
