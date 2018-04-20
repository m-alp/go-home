package experiments

import (
	"net/http"
	"log"
	"fmt"
)

func Http() {
	url := "https://httpbin.org/post"
	ct := "application/json"
	body := "{\"form\": {\"custname\": \"aaa\"}}"

	rst, err := http.Post(url, ct, body)
	if err != nil {
		log.Println("Error!", err)
	}

	fmt.Println("-- Response --")
	fmt.Println("Status Code:", rst.StatusCode)
	fmt.Println("Body:", rst.Body)
}
