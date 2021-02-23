package gomarkdown

import "fmt"

// paragraphConv ...
func (convData *convertedData) paragraphConv() {
	// inline
	inline := convData.inlineConv()

	// if inline or no
	if !inline && !convData.typeChenged {
		convData.lineType = typeNone
		convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	} else if !inline {
		convData.lineType = typeNone
	} else if inline && convData.typeChenged {
		convData.markdownLines[0] = fmt.Sprintf("<p>%s", convData.markdownLines[0])
	}
}

// paragraphClose ...
func (convData *convertedData) paragraphClose() {
	convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
}
