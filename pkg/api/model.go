package api

type requestPayload struct {
	Uuid    string  `json:"uuid" binding:"required"`
	Payload Payload `json:"payload" binding:"required"`
}

type Payload struct {
	Name    string `json:"name" binding:"required"`
	Timeout int    `json:"timeout" binding:"required"`
}
