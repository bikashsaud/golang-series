package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"_"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Decode json data to normal while get")
	DecodeJSON()
}

func DecodeJSON() {
	jsonDataFromUi := []byte(`
	{
		"course_name": "Django",
		"price": 299,
		"platform": "Youtube",
		"tags": ["programming", "backend", "django"]
	}
	`)

	var addCourse course
	checkCourseValid := json.Valid(jsonDataFromUi)

	if checkCourseValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromUi, &addCourse)
		fmt.Printf("%#v\n", addCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	fmt.Printf("data in key value pair... \n")
	// when we need data in key value format
	var newCourseData map[string]interface{} // when create map for online json first is key and value is not gaurantee so make interface,
	json.Unmarshal(jsonDataFromUi, &newCourseData)
	fmt.Printf("%#v\n", newCourseData)

	fmt.Printf("data get format:")
	for k, v := range newCourseData {
		fmt.Printf("key is %v, value is %v and type is %T\n", k, v, v)
	}
}
