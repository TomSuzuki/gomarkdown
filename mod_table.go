package gomarkdown

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// isTableHead ...count "|"
func (convData *convertedData) isTableHead() bool {
	var line = convData.markdownLines[0]
	return (strings.Trim(line, " ") + " ")[:1] == "|" && strings.Count(line, "|") > 1
}

// tableHeadConv ...make align
func (convData *convertedData) tableHeadConv() {
	// error check
	if len(convData.markdownLines) == 1 {
		return
	}

	// align
	alignLine := strings.Split(convData.markdownLines[1], "|")
	convData.markdownLines = append([]string{convData.markdownLines[0]}, convData.markdownLines[2:]...)
	alignLine = alignLine[1 : len(alignLine)-1]
	for _, v := range alignLine {
		convData.tableAlign = append(convData.tableAlign, tableAlign[[2]bool{string(v[len(v)-1]) == ":", string(v[0]) == ":"}])
	}

	// <tr>
	convData.tableGenerate("th")

	// open <table><thead>
	convData.markdownLines[0] = fmt.Sprintf("<table><thead>%s", convData.markdownLines[0])
}

// tableHeadClose ...if table is close
func (convData *convertedData) tableHeadClose() {
	if convData.lineType != typeTableBody {
		convData.shiftLine()
		convData.markdownLines[0] = "</thead></table>"
		convData.tableAlign = nil
	}
}

// ------------------------------------------------------------------------------------------------

// isTableBody ...thead and before type check
func (convData *convertedData) isTableBody() bool {
	return convData.isTableHead() && (convData.lineType == typeTableHead || convData.lineType == typeTableBody)
}

// tableBodyConv ...table generation
func (convData *convertedData) tableBodyConv() {
	// <tr>
	convData.tableGenerate("td")

	// <tbody>
	if convData.typeChenged {
		convData.markdownLines[0] = fmt.Sprintf("</thead><tbody>%s", convData.markdownLines[0])
	}

	// inline
	convData.inlineConv()
}

// tableBodyClose ...
func (convData *convertedData) tableBodyClose() {
	convData.shiftLine()
	convData.markdownLines[0] = "</tbody></table>"
	convData.tableAlign = nil
}

// tableGenerate ...<tr>
func (convData *convertedData) tableGenerate(tagType string) {
	// <tr>
	var reg = `\|` + strings.Repeat(`([^|]*)\|`, len(convData.tableAlign)) + `$`
	var htm string
	for i := range convData.tableAlign {
		htm += fmt.Sprintf("<%s align='%s'>$%s</%s>", tagType, convData.tableAlign[i], strconv.Itoa(i+1), tagType)
	}
	convData.markdownLines[0] = fmt.Sprintf("<tr>%s</tr>", regexp.MustCompile(reg).ReplaceAllString(convData.markdownLines[0], htm))
}

// tableAlign ... : --- :
var tableAlign = map[[2]bool]string{
	{true, false}:  "right",  // : ---
	{false, true}:  "left",   //   --- :
	{true, true}:   "center", // : --- :
	{false, false}: "center", //   ---
}
