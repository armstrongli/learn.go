package dockertool

import (
	"context"
	"testing"

	"github.com/docker/docker/client"
)

func TestRun(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		t.Fatal(err)
	}
	ip, err := Run(context.TODO(), cli, map[string]string{"foo": "bar"}, "alpine:3.13", []string{
		"sh", "-c", "while true; do date ; sleep 2 ; done",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}
