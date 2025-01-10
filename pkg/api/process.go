package api

import (
	"encoding/json"
	"log"
	"os"
)

func process(ch chan string, p payload) {

	// ctx := context.Background()
	// inputQueue := os.Getenv("SQS_SOURCE_QUEUE")
	outputQueue := os.Getenv("SQS_DESTINATION_QUEUE")
	// sqs := awsSdk.NewSqsCoordinator()
	d := Data{
		Info:        "test",
		OutputQueue: outputQueue,
	}
	m := Message{
		Data: d,
	}
	out, err := json.Marshal(m)
	if err != nil {
		log.Fatal("Failed")
	}
	ch <- string(out)
}
