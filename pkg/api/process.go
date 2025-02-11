package api

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
	"github.com/meero-com/hmb-proxy/pkg/pollers"
)

func process(ch chan string, p requestPayload) {
	ddb := aws.NewDdbCoordinator()
	uuid := uuid.NewString()
	requestTable := config.GetConfig("ddb.request_table").(string)

	ddbPayload := aws.DdbPayload{
		Name: p.Payload.Name,
	}
	ddbi := aws.DdbItem{
		Uuid:    p.Uuid,
		Payload: ddbPayload,
	}

	_, err := ddb.Put(context.Background(), requestTable, ddbi)

	if err != nil {
		log.Fatalf("Failed to put item %s into table %s", uuid, requestTable)
	}

	pollers.PollDdb(ch, p.Uuid, ddb)
}
