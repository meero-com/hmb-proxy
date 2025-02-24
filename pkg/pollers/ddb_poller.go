package pollers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/meero-com/hmb-proxy/pkg/aws"
	"github.com/meero-com/hmb-proxy/pkg/config"
)

const (
	pollInterval = 6 * time.Second
)

type DdbPoller struct {
	Ddb aws.DdbCoordinator
}

func (p *DdbPoller) Poll(ch chan string, uuid string) {
	ddb := p.Ddb
	responseTable := config.GetConfig("ddb.response_table").(string)
	ctx := context.Background()

	log.Printf("Trying to retrieve item: %s in table %s", uuid, responseTable)
	for {
		item, err := ddb.Get(ctx, uuid, responseTable)

		if err != nil {
			log.Printf("Failed to retrieve item with Uuid: %s, continue polling", uuid)
			continue
		}

		if item.Item != nil {
			i := aws.DdbItem{}
			err = attributevalue.UnmarshalMap(item.Item, &i)
			if err != nil {
				log.Fatalf("Failed to UnmarshalMap ddb item: %s with model: %s", item.Item, i)
			}
			si, err := json.Marshal(i)
			if err != nil {
				log.Fatalf("Failed to Marshal item: %s", i)
			}
			ch <- string(si)
			close(ch)
			return
		}

		time.Sleep(pollInterval)
	}
}
