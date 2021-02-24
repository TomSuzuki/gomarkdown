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
	var tag = "td"
	var text = ""

	// align
	if convData.tableAlign == nil {
		tag = "th"
		text = "<table>"
		if len(convData.markdownLines) == 1 {
			return
		}
		alignLine := strings.Split(convData.markdownLines[1], "|")
		convData.markdownLines = append([]string{convData.markdownLines[0]}, convData.markdownLines[2:]...)
		alignLine = alignLine[1 : len(alignLine)-1]
		for i := range alignLine {
			switch [2]bool{string(alignLine[i][len(alignLine[i])-1]) == ":", string(alignLine[i][0]) == ":"} {
			case [2]bool{true, false}:
				convData.tableAlign = append(convData.tableAlign, "right")
			case [2]bool{false, true}:
				convData.tableAlign = append(convData.tableAlign, "left")
			default:
				convData.tableAlign = append(convData.tableAlign, "center")
			}
		}
	}

	// <tr>
	var reg = `\|` + strings.Repeat(`([^|]*)\|`, len(convData.tableAlign)) + `$`
	var htm = ""
	for i := range convData.tableAlign {
		htm += fmt.Sprintf("<%s align='%s'>$%s</%s>", tag, convData.tableAlign[i], strconv.Itoa(i+1), tag)
	}
	convData.markdownLines[0] = fmt.Sprintf("%s<tr>%s</tr>", text, regexp.MustCompile(reg).ReplaceAllString(convData.markdownLines[0], htm))

	// inline
	convData.inlineConv()
}

// tableClose ...
func (convData *convertedData) tableClose() {
	convData.shiftLine()
	convData.markdownLines[0] = "</table>"
	convData.tableAlign = nil
}
