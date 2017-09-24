package main

import (
	"fmt"
	"io/ioutil"
)

func printVersion() string {
	result, err := ioutil.ReadFile("./VERSION")
	if err != nil {
		fmt.Println("This application is not versioned")
	}

	content := string(result)
	fmt.Println(content)

	return content
}
