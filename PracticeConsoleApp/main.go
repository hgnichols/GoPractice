package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("myapp.log")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Errors in error log:")
	dataAsString := string(data)
	eachLine := strings.Split(dataAsString, "\n")
	for _, line := range eachLine {
		if strings.Contains(line, "ERROR") {
			fmt.Println(line)
		}
	}
}
