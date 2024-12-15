package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	apis map[int]string
	c    chan map[int]interface{} // channel to store map[int]interface{}
	wg   sync.WaitGroup
)

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}
			json.Unmarshal(body, &result)

			var re = make(map[int]interface{})

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 1 stored")
			case 2:
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
				// store the result into the channel
				c <- re
				fmt.Println("Result for API 2 stored")
			case 3:
				if _, check := result["articles"]; check == true {
					for _, article := range result["articles"].([]interface{}) {
						fmt.Println("News Source:", article.(map[string]interface{})["source"].(map[string]interface{})["name"])
						fmt.Println("News Title:", article.(map[string]interface{})["title"])
						fmt.Println("News Description:", article.(map[string]interface{})["description"])
						fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
					}
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				fmt.Println("Result for exercise printed")
				fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

type Result struct {
	Success   bool
	Timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

type Error struct {
	Success bool
	Error   struct {
		Code int
		Type string
		Info string
	}
}

func main() {
	wg.Add(1)
	c = make(chan map[int]interface{})
	apis = make(map[int]string)
	apis[1] = "http://data.fixer.io/api/latest?access_key=????"
	apis[2] = "http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=????"
	apis[3] = "https://newsapi.org/v2/top-headlines?country=us&category=business&apiKey=????"

	go func() {
		fetchData(1)
	}()
	go func() {
		fetchData(2)
	}()
	go func() {
		defer wg.Done()
		fetchData(3)
	}()

	wg.Wait()

	// we expect two results in the channel
	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Done!")
}
