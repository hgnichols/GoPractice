package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are Debug, Info, and Critical")

	flag.Parse()

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Errors in error log:")
	dataAsString := string(data)
	eachLine := strings.Split(dataAsString, "\n")
	for _, line := range eachLine {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
