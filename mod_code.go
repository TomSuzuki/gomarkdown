package gomarkdown

// isCode ...
func (convData *convertedData) isCode() bool {
	return convData.lineType == typeCodeMarker || convData.lineType == typeCode
}

// isCodeMarker ...
func (convData *convertedData) isCodeMarker() bool {
	return (convData.markdownLines[0] + "   ")[:3] == "```"
}

// codeMarkerConv ...start code lines
func (convData *convertedData) codeMarkerConv() {
	convData.markdownLines[0] = "<pre><code>"
}

// codeClose ...close code lines
func (convData *convertedData) codeClose() {
	if convData.lineType == typeCodeMarker {
		convData.markdownLines[0] = "</code></pre>"
		convData.lineType = typeNone
	}
}
