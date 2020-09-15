package main

import (
	"encoding/json"
	"fmt"

	"github.com/ritabc/twg/stripe/v1"
)

func main() {
	c := stripe.Client{
		Key: "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
	charge, err := c.Charge(2000, "tok_mastercard", "Charge for demo purposes.")
	if err != nil {
		panic(err)
	}
	fmt.Println(charge)
	jsonBytes, err := json.Marshal(charge)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))
}
