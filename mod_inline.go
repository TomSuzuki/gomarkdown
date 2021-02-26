package gomarkdown

import (
	"fmt"
	"regexp"
	"strings"
)

// inlineConv ...replacement in line (regular expressions)
func (convData *convertedData) inlineConv() {

	// inline text
	convData.inlineTage("**", "strong")
	convData.inlineTage("`", "code")
	convData.inlineTage("~", "s")
	convData.inlineTage("__", "em")
	convData.inlineTage("_", "em")
	convData.inlineTage("*", "em")

	// <img> and <a>
	convData.markdownLines[0] = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`).ReplaceAllString(convData.markdownLines[0], "<img alt='$1' src='$2'>")
	convData.markdownLines[0] = regexp.MustCompile(`\[(.*)\]\((.*)\)`).ReplaceAllString(convData.markdownLines[0], "<a href='$2'>$1</a>")

	// <br>
	convData.markdownLines[0] = strings.Replace(convData.markdownLines[0], "  ", "<br>", -1)
}

// inlineTage ...md -> html
func (convData *convertedData) inlineTage(md string, html string) {
	var codeList = strings.Split(convData.markdownLines[0], md)
	convData.markdownLines[0] = ""
	var isEven = len(codeList)%2 == 0

	// insert tags
	for i, v := range codeList {
		if isEven && i == len(codeList)-1 {
			convData.markdownLines[0] += (md + v)
		} else if i%2 == 0 {
			convData.markdownLines[0] += v
		} else if isNotBrokenHTML(v) {
			convData.markdownLines[0] += fmt.Sprintf("<%s>%s</%s>", html, v, html)
		} else {
			convData.markdownLines[0] += (md + v)
		}
	}
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
