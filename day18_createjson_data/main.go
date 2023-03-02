package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Json data create")
	EncodeJsonData()
}

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"_"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJsonData() {

	myCourse := []course{
		{"Reactjs", 299, "Youtube", "admin123", []string{"web", "dev", "reacts"}},
		{"Django", 299, "Youtube", "admin123", []string{"programming", "backend", "django"}},
		{"Python", 299, "Youtube", "admin123", nil},
		{"AI", 299, "Youtube", "admin123", []string{}},
	}

	// package this data as JSON data

	jsonData, err := json.Marshal(myCourse)
	if err != nil {
		panic("failed to marsh")

	}

	fmt.Printf("%s\n", jsonData)
	// marsalindent

	jsonDataIndent, err := json.MarshalIndent(myCourse, "", "\t")

	fmt.Printf(string(jsonDataIndent))

}
