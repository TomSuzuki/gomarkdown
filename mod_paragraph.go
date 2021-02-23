package gomarkdown

import "fmt"

// paragraphConv ...
func paragraphConv(convData convertedData) convertedData {
	// inline
	var inline bool
	convData, inline = inlineConv(convData)

	// if inline or no
	if !inline && !convData.typeChenged {
		convData.lineType = typeNone
		convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	} else if !inline {
		convData.lineType = typeNone
	} else if inline && convData.typeChenged {
		convData.markdownLines[0] = fmt.Sprintf("<p>%s", convData.markdownLines[0])
	}

	return convData
}

// paragraphClose ...
func paragraphClose(convData convertedData) convertedData {
	convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	return convData
}
