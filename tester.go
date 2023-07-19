package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/ads"
	method := "POST"

	payload := strings.NewReader(`{
    "title": "t1",
    "description": "d1",
    "price": 1.1,
    "author": {
        "firstName": "Muka",
        "secondName": "Lamejor Detodas",
        "age": 10,
        "email": "muka@gmail.com"
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic bXVrYTpsYW1lam9yZGV0b2Rhcw==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
