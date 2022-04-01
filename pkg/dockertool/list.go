package dockertool

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func hasAll(left, right map[string]string) bool {
	for k, v := range right {
		if left[k] != v {
			return false
		}
	}
	return true
}

func List(ctx context.Context, cli client.APIClient, labels map[string]string) (containerIds []string, errOut error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	for _, item := range containers {
		if hasAll(item.Labels, labels) && item.State == "running" {
			containerIds = append(containerIds, item.ID)
		}
	}
	return
}
