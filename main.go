package main

import (
	"context"

	sw "splitwise-to-ynab/splitwise"

	"github.com/apex/log"
	"github.com/joho/godotenv"
)

type Expense struct {
	Cost           string `json:"cost"`
	Description    string `json:"description"`
	Details        string `json:"details"`
	Date           string `json:"date"`
	RepeatInterval string `json:"repeat_interval"`
	CurrencyCode   string `json:"currency_code"`
	CategoryId     uint32 `json:"category_id"`
	GroupId        uint32 `json:"group_id"`
}
type ExpenseSplitEqually struct {
	Expense
	SplitEqually bool `json:"split_equally"`
}

func main() {

	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	splitwiseClient := sw.InitClient()

}
