package splitwise

import (
	"os"

	"github.com/anvari1313/splitwise.go"
	"github.com/apex/log"
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
