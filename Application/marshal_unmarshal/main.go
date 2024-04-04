package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

func main() {

	//marshal- converting go struct data into json
	group := ColorGroup{
		Id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	fmt.Println(string(b))
	os.Stdout.Write(b)

	fmt.Println()

	//unmarshal - converting json to go data struct format
	var jsonBlob = []byte(`[{"Name":"Platypus","Order":"Platypus"},{"Name":"Lion","Order":"Lion"}]`)
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		log.Fatalf("error %v ", err)
	}
	fmt.Println(animals)

}
