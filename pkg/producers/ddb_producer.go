package producers

import (
	"context"
	"log"

	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
)

type DdbProducer struct {
	Ddb aws.DdbCoordinator
}

func (p *DdbProducer) Produce(uuid string, payload string) {
	ddb := p.Ddb
	requestTable := config.GetConfig("ddb.request_table").(string)
	ddbPayload := aws.DdbPayload{
		Name: payload,
	}
	ddbi := aws.DdbItem{
		Uuid:    uuid,
		Payload: ddbPayload,
	}

	_, err := ddb.Put(context.Background(), requestTable, ddbi)

	if err != nil {
		log.Fatalf("Failed to put item %s into table %s", uuid, requestTable)
	}
}
