package gomarkdown

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// isTable
func (convData *convertedData) isTable() bool {
	var line = convData.markdownLines[0]
	return (strings.Trim(line, " ") + " ")[:1] == "|" && strings.Count(line, "|") > 1
}

// tableConv ...table generation
func (convData *convertedData) tableConv() {
	var tagType = "td"
	var tagOpen = ""

	// align
	if convData.tableAlign == nil {
		tagType = "th"
		tagOpen = "<table><thead>"
		if len(convData.markdownLines) == 1 {
			return
		}
		alignLine := strings.Split(convData.markdownLines[1], "|")
		convData.markdownLines = append([]string{convData.markdownLines[0]}, convData.markdownLines[2:]...)
		alignLine = alignLine[1 : len(alignLine)-1]
		for _, v := range alignLine {
			convData.tableAlign = append(convData.tableAlign, tableAlign[[2]bool{string(v[len(v)-1]) == ":", string(v[0]) == ":"}])
		}
	}

	// <tr>
	var reg = `\|` + strings.Repeat(`([^|]*)\|`, len(convData.tableAlign)) + `$`
	var htm string
	for i := range convData.tableAlign {
		htm += fmt.Sprintf("<%s align='%s'>$%s</%s>", tagType, convData.tableAlign[i], strconv.Itoa(i+1), tagType)
	}
	convData.markdownLines[0] = fmt.Sprintf("%s<tr>%s</tr>", tagOpen, regexp.MustCompile(reg).ReplaceAllString(convData.markdownLines[0], htm))

	// thead and tbody
	if tagType == "th" {
		convData.markdownLines[0] = fmt.Sprintf("%s</thead><tbody>", convData.markdownLines[0])
	}

	// inline
	convData.inlineConv()
}

// tableClose ...
func (convData *convertedData) tableClose() {
	convData.shiftLine()
	convData.markdownLines[0] = "</tbody></table>"
	convData.tableAlign = nil
}

// tableAlign ... : --- :
var tableAlign = map[[2]bool]string{
	{true, false}:  "right",  // : ---
	{false, true}:  "left",   //   --- :
	{true, true}:   "center", // : --- :
	{false, false}: "center", //   ---
}
