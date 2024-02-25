package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "bond",
		Sayings: []string{"Shaken", "Any last wishes?", "Never say never"},
	}

	bs, err := toJson(p1)
	if err!=nil{
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func toJson(a interface{} )([]byte,error){
	bs,err :=json.Marshal(a) 

	if err!=nil{
		// return []byte{},fmt.Errorf("There was an error in json %v",err)
		return []byte{},errors.New(fmt.Sprintf("There was an error in json %v",err))
	}
	return bs,nil
}