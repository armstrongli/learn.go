package dockertool

import (
	"context"
	"testing"

	"github.com/docker/docker/client"
)

func TestList(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatal(err)
	}
	ids, err := List(context.TODO(), cli, map[string]string{"foo": "bar"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ids)
}
