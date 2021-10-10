package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Handling Json Data in Go Lang")
	// EncodeJson()

	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"React Js BootCamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "node", "mern", "js"},
		},
		{
			"Angular BootCamp", 299, "LearnCodeOnline.in", "abcbcd123", nil,
		},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
        {
                "coursename": "React Js BootCamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDatafromWeb)
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDatafromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON data isn not valid")
	}

	// some cases where you just want to add data to key value pairs

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is:  %T\n", k, v, v)
	}
}
