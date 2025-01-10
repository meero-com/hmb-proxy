package api

type Message struct {
	Data Data
}

type Data struct {
	Info        string
	OutputQueue string
}
