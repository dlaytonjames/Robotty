package robot

import "strings"

type Line struct {
	directive string
	value string
}

/**
	Checks if a line is a user agent directive
 */
func (line *Line) IsUserAgent() bool {
	if strings.ToLower(strings.TrimSpace(line.directive)) == "user-agent" {
		return true
	}
	return false
}

func (line *Line) IsDisallow() bool {
	if strings.ToLower(strings.TrimSpace(line.directive)) == "disallow" {
		return true
	}
	return false
}



