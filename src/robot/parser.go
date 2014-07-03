package robot

import (
	"strings"
	"net/http"
	"io/ioutil"
	"errors"
	"fmt"
)



/**
	Get all robots.txt lines and create a
	decision object
 */
func FromString(content string) *Decision{

	// split content and get all lines
	lines := getLines(content)

	// pass and return a decision
	return &Decision{ lines: lines }
}

/**
	Get all robots.txt lines from a http.Response object
	and then create a decision object
 */
func FromResponse(response *http.Response) (*Decision, error) {

	defer func(){
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// read from response
	if resByte, err := ioutil.ReadAll(response.Body); err == nil {
		lines := getLines(string(resByte))
		return &Decision{ lines: lines }, nil
	} else {
		return new(Decision), errors.New("unable to read response object")
	}
}

/**
	Get line objects from robots.txt content
 */
func getLines(content string) []*Line {

	// holds line objects
	lineObjs := []*Line{}

	// split into separate lines
	lines := strings.Split(content, "\n")

	// convert each line to line objects
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			lineParts := strings.Split(line, ":")
			lineObjs = append(lineObjs, &Line{ directive: lineParts[0], value: strings.TrimSpace(lineParts[1]) })
		}
	}
	return lineObjs
}

