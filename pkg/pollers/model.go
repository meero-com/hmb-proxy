package pollers

import "time"

type Poller interface {
	Poll(ch chan string, uuid string)
}

const (
	pollInterval = 6 * time.Second
)
