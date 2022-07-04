package main

import (
	"context"
	"os"
	"strconv"

	"github.com/anvari1313/splitwise.go"
	"github.com/apex/log"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var splitwiseApiKey string
	if val, ok := os.LookupEnv("SPLITWISE_API_KEY"); ok {
		splitwiseApiKey = val
	} else {
		os.Exit(1)
	}

	auth := splitwise.NewAPIKeyAuth(splitwiseApiKey)
	client := splitwise.NewClient(auth)
	ctx := context.Background()

	groups, err := client.Groups(ctx)
	if err != nil {
		log.Error(err.Error())
	}
	for _, v := range groups {
		if v.Name == "Appartamento" {
			id := strconv.FormatUint(v.ID, 10)
			log.Info(id)
		}

	}
}
