package gomarkdown

// isCode ...
func (convData *convertedData) isCode() bool {
	return convData.lineType == typeCodeMarker || convData.lineType == typeCode
}

// isCodeMarker ...
func (convData *convertedData) isCodeMarker() bool {
	return len(convData.markdownLines[0]) >= 3 && (convData.markdownLines[0])[:3] == "```"
}

// convCodeMarker ...start code lines
func (convData *convertedData) convCodeMarker() {
	var text = ""

	if len(convData.markdownLines) >= 2 {
		convData.markdownLines = convData.markdownLines[1:]
		if !convData.isCodeMarker() {
			text = convData.markdownLines[0]
		} else {
			convData.markdownLines = append(convData.markdownLines[:1], convData.markdownLines...)
			convData.markdownLines[1] = "```"
		}
	}

	convData.markdownLines[0] = "<pre><code>" + text

	convData.lineType = typeCode
}

// closeCode ...close code lines
func (convData *convertedData) closeCode() {
	if convData.lineType == typeCodeMarker {
		convData.markdownLines[0] = "</code></pre>"
		convData.lineType = typeNone
	}
}
