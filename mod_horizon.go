package gomarkdown

import (
	"strings"
)

// isHorizon
func (convData *convertedData) isHorizon() bool {
	var line = strings.Replace(convData.markdownLines[0], " ", "", -1)
	return len(line) >= 3 && (len(line) == strings.Count(line, "-") || len(line) == strings.Count(line, "_") || len(line) == strings.Count(line, "*"))
}

// convHorizon
func (convData *convertedData) convHorizon() {
	convData.markdownLines[0] = "<hr>"
}
