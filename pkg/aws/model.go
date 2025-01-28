package aws

type DdbItem struct {
	Uuid    string     `json:"uuid"`
	Payload DdbPayload `json:"payload"`
}

type DdbPayload struct {
	Name string `json:"name"`
}
