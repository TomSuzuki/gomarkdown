package gomarkdown

import (
	"strings"
)

// isTableHead ...count "|"
func (convData *convertedData) isTableHead() bool {
	var line = convData.markdownLines[0]
	return len(line) >= 2 && strings.Trim(line, " ")[:1] == "|" && strings.Count(line, "|") > 1
}

// convTableHead ...make align
func (convData *convertedData) convTableHead() {
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
	convData.markdownLines[0] = ("<table><thead>" + convData.markdownLines[0])
}

// closeTableHead ...if table is close
func (convData *convertedData) closeTableHead() {
	if convData.lineType != typeTableBody {
		convData.shiftLine()
		convData.markdownLines[0] = "</thead></table>"
		convData.tableAlign = nil
	}
}

// ------------------------------------------------------------------------------------------------

// isTableBody ...thead and before type check
func (convData *convertedData) isTableBody() bool {
	return (convData.lineType == typeTableBody || convData.lineType == typeTableHead) && convData.isTableHead()
}

// convTableBody ...table generation
func (convData *convertedData) convTableBody() {
	// <tr>
	convData.tableGenerate("td")

	// <tbody>
	if convData.typeChenged {
		convData.markdownLines[0] = ("</thead><tbody>" + convData.markdownLines[0])
	}

	// inline
	convData.inlineConv()
}

// closeTableBody ...
func (convData *convertedData) closeTableBody() {
	convData.shiftLine()
	convData.markdownLines[0] = "</tbody></table>"
	convData.tableAlign = nil
}

// tableGenerate ...<tr>
func (convData *convertedData) tableGenerate(tagType string) {
	// check
	var tr = strings.Split(convData.markdownLines[0], "|")
	if len(tr)-2 <= 1 || len(tr)-2 != len(convData.tableAlign) || tr[0] != "" || tr[len(tr)-1] != "" {
		return
	}

	// make
	var html []string
	html = append(html, "<tr>")
	for i, v := range convData.tableAlign {
		html = append(html, "<")
		html = append(html, tagType)
		html = append(html, " align='")
		html = append(html, v)
		html = append(html, "'>")
		html = append(html, tr[i+1])
		html = append(html, "</")
		html = append(html, tagType)
		html = append(html, ">")
	}
	html = append(html, "</tr>")
	convData.markdownLines[0] = strings.Join(html, "")
}

// tableAlign ... : --- :
var tableAlign = map[[2]bool]string{
	{true, false}:  "right",  // : ---
	{false, true}:  "left",   //   --- :
	{true, true}:   "center", // : --- :
	{false, false}: "center", //   ---
}
