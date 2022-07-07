package splitwise

import (
	"context"
	"encoding/json"
	"os"

	"github.com/apex/log"
	"github.com/rtrive/splitwise.go"
)

func InitClient() splitwise.Client {
	var splitwiseApiKey string
	if val, ok := os.LookupEnv("SPLITWISE_API_KEY"); ok {
		splitwiseApiKey = val
	} else {
		log.Error("No splitwise api key present")
		os.Exit(1)
	}

	log.Debug("Init Client")
	auth := splitwise.NewAPIKeyAuth(splitwiseApiKey)
	client := splitwise.NewClient(auth)
	return client
}

func CreateExpenseSplitEqually(ctx context.Context, sw splitwise.Client, transaction string) {
	var txSplitEqually splitwise.ExpenseSplitEqually
	err := json.Unmarshal([]byte(transaction), &txSplitEqually)
	if err != nil {
		log.Error("wrong conversion")
	}

	_, err = sw.CreateExpenseSplitEqually(ctx, txSplitEqually)

	if err != nil {
		log.Error("Error during create transaction " + err.Error())
	}
}
