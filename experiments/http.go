package experiments

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
)

func Http() {
	//get()
	post()
}

func get() {
	u := "https://httpbin.org/get"

	rsp, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	fmt.Println("get:\n", string(body))
}


func post() {
	u := "https://httpbin.org/post"
	//ct := "application/json"
	//body := "{\"form\": {\"custname\": \"aaa\"}}"

	rsp, err := http.PostForm(u, url.Values{"q": {"custname"}})
	if err != nil {
		panic(err)
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	fmt.Println("post:\n", string(body))
}
