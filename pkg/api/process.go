package api

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/meero-com/guild-proxy/pkg/aws"
	"github.com/meero-com/guild-proxy/pkg/config"
	"github.com/meero-com/guild-proxy/pkg/pollers"
)

func process(ch chan string, p payload) {
	ddb := aws.NewDdbCoordinator()
	uuid := uuid.NewString()
	//responseTable := config.GetConfig("ddb.response_table").(string)
	requestTable := config.GetConfig("ddb.request_table").(string)

	// ctx := context.Background()
	// inputQueue := config.GetConfig("sqs.input_queue").(string)
	// sqs := awsSdk.NewSqsCoordinator()
	ddbPayload := aws.DdbPayload{
		Name: p.Name,
	}
	ddbi := aws.DdbItem{
		Uuid:    uuid,
		Payload: ddbPayload,
	}

	_, err := ddb.Put(context.Background(), requestTable, ddbi)

	if err != nil {
		log.Fatal("Failed to put item %s into table %s", uuid, requestTable)
	}

	//mockService(uuid, responseTable, ddb)
	i, err := pollers.PollDdb(uuid, ddb)

	if err != nil {
		log.Fatal("Failed to poll item: %s from ddb", uuid)
	}

	ch <- i
}

// Used to Mock external service
func mockService(uuid string, table string, ddb aws.DdbCoordinator) {
	ddbPayload := aws.DdbPayload{
		Name: "Success!",
	}
	ddbi := aws.DdbItem{
		Uuid:    uuid,
		Payload: ddbPayload,
	}
	_, err := ddb.Put(context.Background(), table, ddbi)

	if err != nil {
		log.Fatal("mockService failed to put item in ddb")
	}
}
