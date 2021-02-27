package gomarkdown

import (
	"strings"
)

// isHorizon
func (convData *convertedData) isHorizon() bool {
	var line = strings.Replace(convData.markdownLines[0], " ", "", -1)
	return len(line) >= 3 && (line[:3] == "---" || line[:3] == "___" || line[:3] == "***")
}

// convHorizon
func (convData *convertedData) convHorizon() {
	convData.markdownLines[0] = "<hr>"
}
