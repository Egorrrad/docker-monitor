package model

import "time"

type ContainerStatus struct {
	IP          string    `json:"ip"`
	Alive       bool      `json:"alive"`
	Checked     time.Time `json:"checked"`
	LastSuccess time.Time `json:"lastSuccess"`
}
