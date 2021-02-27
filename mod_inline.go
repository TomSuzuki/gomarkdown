package gomarkdown

import (
	"strings"
)

// inlineConv ...replacement in line (regular expressions)
func (convData *convertedData) inlineConv() {

	// inline text
	convData.inlineTag("**", "strong")
	convData.inlineTag("`", "code")
	convData.inlineTag("~", "s")
	convData.inlineTag("__", "em")
	convData.inlineTag("_", "em")
	convData.inlineTag("*", "em")

	// <img> and <a>
	convData.inlineLink("![", "](", ")", "<img alt='$1' src='$2'>")
	convData.inlineLink("[", "](", ")", "<a href='$2'>$1</a>")

	// <br>
	convData.markdownLines[0] = strings.Replace(convData.markdownLines[0], "  ", "<br>", -1)
}

// inlineLink ...![]() and []()
func (convData *convertedData) inlineLink(start, middle, end string, format string) {
	var line = convData.markdownLines[0]
	var p = 0

	for true {
		// index
		var iMiddle = p + strings.Index(line[p:], middle)
		if iMiddle == -1 {
			break
		}
		var iStart = strings.LastIndex(line[:iMiddle], start)
		var iEnd = iMiddle + strings.Index(line[iMiddle:], end)

		// replace
		if iMiddle != -1 && iStart != -1 && iEnd != -1 {
			var s = strings.NewReplacer(
				"$1", line[len(start)+iStart:iMiddle],
				"$2", line[len(middle)+iMiddle:iEnd],
			).Replace(format)
			line = line[:iStart] + s + line[len(end)+iEnd:]
		}

		// error check
		p = iMiddle + len(middle)
		if p > len(line) {
			break
		}
	}

	convData.markdownLines[0] = line
}

// indexList
func indexList(s, substr string) []int {
	var n []int
	for true {
		var m = strings.Index(s, substr)
		if m == -1 {
			break
		}
		n = append(n, m)
		s = s[m:]
	}
	return n
}

// inlineTag ...md -> html
func (convData *convertedData) inlineTag(md string, html string) {
	var codeList = strings.Split(convData.markdownLines[0], md)
	var isEven = len(codeList)%2 == 0
	var text []string
	convData.markdownLines[0] = ""

	// insert tags
	for i, v := range codeList {
		if isEven && i == len(codeList)-1 {
			text = append(text, md)
			text = append(text, v)
		} else if i%2 == 0 {
			text = append(text, v)
		} else if isNotBrokenHTML(v) {
			text = append(text, "<")
			text = append(text, html)
			text = append(text, ">")
			text = append(text, v)
			text = append(text, "</")
			text = append(text, html)
			text = append(text, ">")
		} else {
			text = append(text, md)
			text = append(text, v)
		}
	}

	// join
	convData.markdownLines[0] = strings.Join(text, "")
}

// isNotBrokenHTML ..."<s></s><em></em>" <<< true, "<s><em></em>" <<< false, "<img ...>" <<< false...?
func isNotBrokenHTML(html string) bool {
	var nest = 0
	var open = false
	for i := 0; i < len(html); i++ {
		s := html[i]
		if open {
			if string(s) == "/" {
				nest--
			} else {
				nest++
			}
		}
		open = string(s) == "<"
		if nest == -1 {
			return false
		}
	}
	return nest == 0
}
