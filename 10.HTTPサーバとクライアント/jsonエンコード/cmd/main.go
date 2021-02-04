package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Person is person
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{
		Name: "iciro",
		Age:  100,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
