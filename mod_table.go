package gomarkdown

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
		if len(alignLine) < 2 {
			return
		}
		if len(convData.markdownLines) > 2 {
			convData.markdownLines = append([]string{convData.markdownLines[0]}, convData.markdownLines[2:]...)
		} else {
			convData.markdownLines = []string{convData.markdownLines[0]}
		}
		alignLine = alignLine[1 : len(alignLine)-1]
		for i := range alignLine {
			colonR := string(alignLine[i][len(alignLine[i])-1]) == ":"
			colonL := string(alignLine[i][0]) == ":"
			if colonR && !colonL {
				convData.tableAlign = append(convData.tableAlign, "right")
			} else if !colonR && colonL {
				convData.tableAlign = append(convData.tableAlign, "left")
			} else {
				convData.tableAlign = append(convData.tableAlign, "center")
			}
		}
	}

	// <tr>
	var reg = `\|`
	var htm = ""
	for i := range convData.tableAlign {
		reg += `([^|]*)\|`
		htm += fmt.Sprintf("<%s align='%s'>$%s</%s>", tag, convData.tableAlign[i], strconv.Itoa(i+1), tag)
	}
	reg += `$`
	text += "<tr>" + regexp.MustCompile(reg).ReplaceAllString(convData.markdownLines[0], htm) + "</tr>"
	convData.markdownLines[0] = text

	// inline
	convData.inlineConv()
}

// tableClose ...
func (convData *convertedData) tableClose() {
	convData.markdownLines[0] = fmt.Sprintf("</table>%s", convData.markdownLines[0])
	convData.tableAlign = nil
}
