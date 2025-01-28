package pollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/meero-com/guild-proxy/pkg/aws"
	"github.com/meero-com/guild-proxy/pkg/config"
)

func PollDdb(ch chan string, uuid string, ddb aws.DdbCoordinator) {
	responseTable := config.GetConfig("ddb.response_table").(string)
	fmt.Println(responseTable)
	ctx := context.Background()

	log.Println("Trying to retrieve item: %s in table %s", uuid, responseTable)
	item, err := ddb.Get(ctx, uuid, responseTable)
	if err != nil {
		log.Fatalf("Failed to poll item in ddb: %s", responseTable)
	}

	i := aws.DdbItem{}
	err = attributevalue.UnmarshalMap(item.Item, &i)
	if err != nil {
		println("Failed to UnmarshalMap")
	}
	si, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Failed to Marshal item: %s, error: %s", uuid, err)
	}

	fmt.Println(i)
	fmt.Println(item.Item)
	fmt.Println(string(si))
	ch <- string(si)
}
