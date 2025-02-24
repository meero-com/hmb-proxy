package pollers

type Poller interface {
	Poll(ch chan string, uuid string)
}
