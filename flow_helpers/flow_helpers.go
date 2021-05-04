package flow_helpers

import (
	"context"
	"time"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
)

func GetLatestBlockId(ctx context.Context, c *client.Client) (flow.Identifier, error) {
	block, err := c.GetLatestBlock(ctx, true)
	if err != nil {
		return flow.Identifier{}, err
	}
	return block.ID, nil
}

func WaitForSeal(ctx context.Context, c *client.Client, id flow.Identifier) (result *flow.TransactionResult, err error) {
	result, err = c.GetTransactionResult(ctx, id)
	if err != nil {
		return
	}

	if result.Error != nil {
		err = result.Error
		return
	}

	for result.Status != flow.TransactionStatusSealed {
		time.Sleep(time.Second)
		result, err = c.GetTransactionResult(ctx, id)

		if err != nil {
			return
		}

		if result.Error != nil {
			err = result.Error
			return
		}
	}

	return result, nil
}