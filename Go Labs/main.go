package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string // first char must be capitalized
	Lastname  string // first char must be capitalized
	Details   struct {
		Height int
		Weight float32
	}
}

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func main() {
	// var persons []People
	// jsonString :=
	// 	`[
	// 		{
	// 			"firstname":"Wei-Meng",
	// 			"lastname":"Lee",
	// 			"details": {
	// 				"height":175,
	// 				"weight":70.0
	// 			}
	// 		},
	// 		{
	// 			"firstname":"Mickey",
	// 			"lastname":"Mouse",
	// 			"details": {
	// 				"height":105,
	// 				"weight":85.5
	// 			}
	// 		}
	// 	]`

	// err := json.Unmarshal([]byte(jsonString), &persons)
	// for _, v := range persons {
	// 	fmt.Println(v.Firstname)
	// 	fmt.Println(v.Lastname)
	// 	fmt.Println(v.Details.Height)
	// 	fmt.Println(v.Details.Weight)
	// }
	// fmt.Println(err) // <nil>

	// jsonString2 :=
	// 	`{
	// 		"base currency":"EUR",
	// 		"destination currency":"USD"
	// 	}`
	// var rates Rates
	// json.Unmarshal([]byte(jsonString2), &rates)
	// fmt.Println(rates.Base)   // EUR
	// fmt.Println(rates.Symbol) // USD

	jsonString3 :=
		`{
			"success": true,
			"timestamp": 1588779306,
			"base": "EUR",
			"date": "2020-05-06",
			"rates": {
				"AUD": 1.683349,
				"CAD": 1.528643,
				"GBP": 0.874757,
				"SGD": 1.534513,
				"USD": 1.080054
			}
		}`
	var result map[string]interface{}
	/*
	   Each string corresponds to a JSON property, and its mapped
	   interface{} type corresponds to the value, which can be of
	   any type.
	   The type is asserted from this interface{} type as is needed in the code.
	*/
	json.Unmarshal([]byte(jsonString3), &result)
	fmt.Println(result["success"])

	currRates := result["rates"] // value of rates is actually an interface{},
	// which could be anything - a map, a
	// string, or an int.
	fmt.Println(currRates) // map[AUD:1.683349 CAD:1.528643
	// GBP:0.874757 SGD:1.534513 USD:1.080054]
	SGD := currRates.(map[string]interface{})["SGD"] // you need to assert it to
	// a map with expected
	// key/value types
	fmt.Println(SGD)
}
