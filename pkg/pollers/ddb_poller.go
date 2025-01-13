package pollers

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/meero-com/guild-proxy/pkg/aws"
)

func PollDdb(uuid string, ddb aws.DdbCoordinator) (string, error) {
	responseTable := os.Getenv("DDB_RESPONSE_TABLE")
	ctx := context.Background()

	item, err := ddb.Get(ctx, uuid, responseTable)
	if err != nil {
		log.Fatal("Failed to poll item in ddb: %s", responseTable)
	}

	i := aws.DdbItem{}
	err = attributevalue.UnmarshalMap(item.Item, &i)

	si, err := json.Marshal(i)
	if err != nil {
		log.Fatal("An error occured")
	}

	return string(si), nil
}
