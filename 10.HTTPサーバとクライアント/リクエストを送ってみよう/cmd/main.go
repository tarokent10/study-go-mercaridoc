package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/omikuji?p=Gopher", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("response: %s", string(b))
}
