package api

type Message struct {
	Header string
	Data   Data
}

type Data struct {
	Uuid        string
	OutputQueue string
}
