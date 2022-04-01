package models

type RankServiceBackend struct {
	Name     string          `json:"name"`
	Expected ExpectedBackend `json:"expected"`
	Status   StatusBackend   `json:"status"`
}

type ExpectedBackend struct {
	Image   string   `json:"image"`
	Command []string `json:"command"`
	Count   int      `json:"count"`
}

type StatusBackend struct {
	RunningCount int      `json:"runningCount"`
	InstanceIPs  []string `json:"instanceIPs"`
}
