package aws

type DdbItem struct {
	Uuid    string
	Payload DdbPayload
}

type DdbPayload struct {
	Name string
}
