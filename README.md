<div align="right">
    <a href="./README.md">English</a> | <a href="./README_jp.md">日本語</a>
</div>

# gomarkdown
markdown parser for golang

## install
```
go get -u github.com/TomSuzuki/gomarkdown
```

## how to use
```
md, _ = ioutil.ReadFile("./markdown.md")
html := gomarkdown.MarkdownToHTML(string(md))
```
