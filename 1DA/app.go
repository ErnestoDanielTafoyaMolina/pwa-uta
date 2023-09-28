package main

import (
	"fmt"

	md "github.com/ErnestoDanielTafoyaMolina/1DA/Modules"
)

func main() {
	var test1 = make(map[string]string)
	test1["test1"] = "asdf"

	fmt.Println("Hello", md.Name)
}
