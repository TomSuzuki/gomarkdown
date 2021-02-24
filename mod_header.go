package gomarkdown

import (
	"regexp"
	"strings"
)

// isHeader ...markdownLines[0] is header?
func (convData *convertedData) isHeader() bool {
	line := strings.Split(strings.Trim(convData.markdownLines[0], " "), " ")[0]
	return line != "" && strings.Trim(line, "#") == ""
}

// headerConv ...
func (convData *convertedData) headerConv() {
	// <h1> - <h6>
	convData.markdownLines[0] = strings.Trim(convData.markdownLines[0], " ")
	for md, html := range map[string]string{
		`^#\s([^#]*?$)`:      "<h1>$1</h1>",
		`^##\s([^#]*?$)`:     "<h2>$1</h2>",
		`^###\s([^#]*?$)`:    "<h3>$1</h3>",
		`^####\s([^#]*?$)`:   "<h4>$1</h4>",
		`^#####\s([^#]*?$)`:  "<h5>$1</h5>",
		`^######\s([^#]*?$)`: "<h6>$1</h6>",
	} {
		convData.markdownLines[0] = regexp.MustCompile(md).ReplaceAllString(convData.markdownLines[0], html)
	}

	// inline
	convData.inlineConv()
}
