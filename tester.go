package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main2() {
	url := "http://localhost:8080/videos"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
