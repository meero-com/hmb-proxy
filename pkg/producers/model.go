package producers

type Producer interface {
	Produce(uuid string, payload string)
}
