package main

import (
	"context"
	"strconv"

	sw "splitwise-to-ynab/splitwise"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

func main() {

	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	splitwiseClient := sw.InitClient()

	groups, err := splitwiseClient.Groups(ctx)
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
