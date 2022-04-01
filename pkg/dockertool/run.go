package dockertool

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Run(ctx context.Context, cli client.APIClient, labels map[string]string, image string, cmd []string) (ip string, errOut error) {
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:  image,
		Cmd:    cmd,
		Tty:    false,
		Labels: labels,
	}, nil, nil, nil, "")
	if err != nil {
		return "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	details, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return "", err
	}
	return details.NetworkSettings.IPAddress, nil
}
