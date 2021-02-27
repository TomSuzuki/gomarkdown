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
	convData.markdownLines[0] = "<pre><code>"
}

// closeCode ...close code lines
func (convData *convertedData) closeCode() {
	if convData.lineType == typeCodeMarker {
		convData.markdownLines[0] = "</code></pre>"
		convData.lineType = typeNone
	}
}
