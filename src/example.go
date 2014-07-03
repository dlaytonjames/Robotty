package main

import (
	"fmt"
	"robot"
	"os"
	"io/ioutil"
	"path"
	"net/http"
)

func TestFromString(){
	dir, _ := os.Getwd()
	if byt, err := ioutil.ReadFile(path.Join(dir, "misc", "robots.txt")); err == nil {
		if err != nil { return }
		result := robot.FromString(string(byt)).IsAllowed("http://google.com/great/bleh", "*")
		if result {
			fmt.Println("Can visit link")
		}
	} else {
		fmt.Println("File doesn't exist")
	}
}

func TestFromResponse(){
	resp, _ := http.Get("http://google.com.ng/robots.txt")
	decision, err := robot.FromResponse(resp)

	// if not problem with reading from response object
	if decision == nil && err == nil {
		fmt.Println("Unable to read from response")
	} else {
		if decision.IsAllowed("http://google.com.ng/bleh/bleh", "*") {
			fmt.Println("Can visit link")
		} else {
			fmt.Println("Can't visit link")
		}
	}
}

func main() {
	//TestFromString()
	TestFromResponse()
}
