package main

import (
	"fmt"
)

func main() {
	var a map[string]interface{}
	a = make(map[string]interface{})

	var b map[string]string

	b = make(map[string]string)
	b["afaf"] = "afdaf"

	a["adf"] = b

	fmt.Println(a)
}
