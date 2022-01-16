package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config
}

func (ws *WebServerV2) reload() {
	ws.config = &Config{
		Content: fmt.Sprintf("%d", time.Now().UnixNano()),
	}
}

func (ws *WebServerV2) ReloadWorker() {
	for {
		time.Sleep(10 * time.Millisecond)
		ws.reload()
	}
}

func (ws *WebServerV2) Visit() string {
	return ws.config.Content
}
