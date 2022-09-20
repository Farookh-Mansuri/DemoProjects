package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Id         int `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
}

func main() {

	u1 := User{1, "John Doe", "gardener"}

	json_data, err := json.Marshal(u1)
	content, err := json.Marshal(u1)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data))

	users := []User{
		{Id: 2, Name: "Roger Roe", Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith", Occupation: "teacher"},
		{Id: 4, Name: "David Brown", Occupation: "programmer"},
	}

	json_data2, err := json.Marshal(users)
	err2 := ioutil.WriteFile("juserData.json", content, 0644)
	if err2 != nil {

		log.Fatal(err)
	}
	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(string(json_data2))
}
