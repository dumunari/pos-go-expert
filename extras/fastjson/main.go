package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonStr := `{"user": {"name": "John", "age": 30}}`
	var p fastjson.Parser

	v, err := p.Parse(jsonStr)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	userJSON := v.Get("user").String()

	var user User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		fmt.Println("Error unmarshaling user JSON:", err)
		return
	}

	fmt.Println("User:", user)
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}
