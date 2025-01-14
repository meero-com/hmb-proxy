package api

import (
	"encoding/json"
	"log"

	"github.com/meero-com/guild-proxy/pkg/config"
)

func process(ch chan string, p payload) {

	// ctx := context.Background()
	// inputQueue := config.GetConfig("sqs.input_queue").(string)
	outputQueue := config.GetConfig("sqs.output_queue").(string)
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
