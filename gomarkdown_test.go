/*
 * gomarkdown
 * available at https://github.com/TomSuzuki/gomarkdown/
 *
 * Copyright 2021 TomSuzuki
 * LICENSE: MIT
 *
 */

//package gomarkdown
package gomarkdown_test

import (
	"testing"

	"github.com/TomSuzuki/gomarkdown"
)

func Test(t *testing.T) {
	test := `
# Header 1
### Header 3
---
<br>  
**Bold**

- item
- item
- item
- **Bold**
- **bold** *em*

- 11
  - 22
  - 22
    - 33
    - 33
  1. aa
  1. aa
    - 33
    - 33


1. aaa
1. bbb
1. ccc

日本語

- 日本語

![text](/path/a.jpg)

[text](link)

`

	test += "`code`and`code`\n"
	test += "```\n"
	test += "go run main.go\n"
	test += "```\n"

	t.Log(gomarkdown.MarkdownToHTML(test))
}
