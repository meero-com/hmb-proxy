package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/meero-com/guild-proxy/pkg/aws"
	"github.com/meero-com/guild-proxy/pkg/config"
	"github.com/meero-com/guild-proxy/pkg/pollers"
)

func process(ch chan string, p requestPayload) {
	ddb := aws.NewDdbCoordinator()
	uuid := uuid.NewString()
	requestTable := config.GetConfig("ddb.request_table").(string)

	fmt.Println(requestTable)
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

	time.Sleep(5 * time.Second)

	pollers.PollDdb(ch, p.Uuid, ddb)

	if err != nil {
		log.Fatalf("Failed to poll item: %s from ddb", uuid)
	}
}
