package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8000/user/delete/1"
	//url := "http://localhost:8000/user/update/1"
	// url := "http://localhost:8000/user/new"
	//fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"username": "name", "first_name": "first", "last_name": "lastname"}`)

	fmt.Println("URL:>", url)


	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}