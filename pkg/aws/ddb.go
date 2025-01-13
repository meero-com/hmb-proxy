package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DdbCoordinator struct {
	DdbClient *dynamodb.Client
}

func NewDdbCoordinator() DdbCoordinator {
	sdkCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-1"))
	if err != nil {
		log.Fatal("Could not load config: %s\n", err.Error())
	}
	ddbClient := dynamodb.NewFromConfig(sdkCfg)
	return DdbCoordinator{
		DdbClient: ddbClient,
	}
}

func (d DdbCoordinator) Get(c context.Context, uuid string, t string) (*dynamodb.GetItemOutput, error) {
	i, err := d.DdbClient.GetItem(
		c,
		&dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"uuid": &types.AttributeValueMemberS{Value: uuid},
			},
			TableName: aws.String(t),
		},
	)
	if err != nil {
		log.Fatal(err)
		return i, err
	}
	return i, nil
}

func (d DdbCoordinator) Put(c context.Context, t string, ddbi DdbItem) (*dynamodb.PutItemOutput, error) {
	i, err := d.DdbClient.PutItem(c, &dynamodb.PutItemInput{
		Item: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: ddbi.Uuid},
			"payload": &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					"name": &types.AttributeValueMemberS{Value: ddbi.Payload.Name},
				},
			},
		},
		TableName: aws.String(t),
	})
	if err != nil {
		log.Fatal("Failed to create no item into table: %s", t)
		return i, err
	}
	return i, nil
}
