package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine/log"
)

func main() {
	ctx := context.Background()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get("http://google.co.jp")
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
	defer res.Body.Close()
	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
}
