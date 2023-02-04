package postgres

import (
	"context"
	"fmt"
	"gimme/database"
	"gimme/docker"
	"gimme/util"
	"github.com/docker/go-connections/nat"
	"io"
	"net"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const POSTGRES = "postgres"

func Start() docker.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	reader, err := cli.ImagePull(ctx, "postgres", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	username := util.GenerateName()
	password := "password"
	containerName := fmt.Sprintf("gimme-%s-%s", username, POSTGRES)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	err = listener.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting postgres container: %s using port %d with username: %s and password: %s", containerName, port, username, password)
	//docker run -e POSTGRES_PASSWORD=secrect -e POSTGRES_USER=postgres <other options> image/name
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "postgres",
		Env:   []string{fmt.Sprintf("POSTGRES_PASSWORD=%s", password), fmt.Sprintf("POSTGRES_USER=%s", username)},
		Tty:   false,
		ExposedPorts: nat.PortSet{
			"5432": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", port),
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	database := database.Database{
		Port:     port,
		Username: username,
		Password: password,
	}

	return docker.Container{
		Name:     containerName,
		Database: database,
	}
}
