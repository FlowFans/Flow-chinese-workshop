package main

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

func main() {
	flowClient, err := client.New("access.devnet.nodes.onflow.org:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	blockEvents, err := flowClient.GetEventsForHeightRange(context.Background(), client.EventRangeQuery{
		Type:        "A.7e60df042a9c0868.FlowToken.TokensDeposited",
		StartHeight: 30052640,
		EndHeight:   30052680,
	})

	for _, blockEvent := range blockEvents {
		for _, event := range blockEvent.Events {
			sum := event.Value.Fields[0].(cadence.UFix64).ToGoValue().(uint64)
			id := event.Value.Fields[1].(cadence.Optional).Value.(cadence.Address).String()
			fmt.Println("数量--- ", sum)
			fmt.Println("ID -- ", id)
		}
	}
}