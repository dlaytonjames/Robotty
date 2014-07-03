package robot

import (
	"strings"
	"net/url"
	"errors"
	"regexp"
)

type Decision struct {
	lines []*Line
}


/**
	Checks if a user agent exists and has group
 */
func (decision *Decision) HasGroup(userAgent string) bool {
	for _, line := range decision.lines {
		if line.IsUserAgent() && line.value == strings.ToLower(userAgent) {
			return true
		}
	}
	return false
}


/**
	Get all directives in a
 */
func (decision *Decision) GetGroup(userAgent string) []*Line {
	lines := []*Line{}
	groupStart := false

	for _, line := range decision.lines {

		if line.IsUserAgent() {
			if line.value == strings.ToLower(userAgent) {
				groupStart = true
			} else {
				groupStart = false
			}
		}

		if groupStart {
			if !line.IsUserAgent() {
				lines = append(lines, line)
			}
		}
	}
	return lines
}

func (decision *Decision) GetURLPath(dUrl string) (string, error) {
	if u, err := url.Parse(dUrl); err == nil {
		return u.Path, nil
	} else {
		return "", errors.New("invalid url supplied")
	}
}


/**
	Takes a directive path and matches it against a target path. It also
	 normalizes the directive path to enforce expected behavior when used
	 as a regex pattern
 */
func (decision *Decision) Matches(targetPath string, directivePath string) bool{

	// modify directive path
	directivePath = strings.Replace(directivePath, ".", "[.]{1}", -1)
	directivePath = strings.Replace(directivePath, "*", ".*?", -1)

	m, _ := regexp.MatchString(directivePath, targetPath)
	return m
}

/**
	Checks if a url can be visited.
	Returns true if its allowed or false otherwise
 */
func (decision *Decision) IsAllowed(url string, userAgent string) bool {

	if decision == nil { return false }

	// extract url path from url
	if urlPath, err := decision.GetURLPath(url); err == nil {

		// holds the group directives
		group := []*Line{}

		// get group directives
		if decision.HasGroup(userAgent) && userAgent != "*" {
			group = decision.GetGroup(userAgent)
		} else {
			group = decision.GetGroup("*")
		}

		// process the group directives
		for _, line := range group {
			if decision.Matches(urlPath, line.value) {
				if line.IsDisallow() { return false } else { return true }
			}
		}
	}
	return true
}
