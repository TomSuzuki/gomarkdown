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

## how to use test
```
go test -v
```
or
```
gotest -v
```
or
```
gotest -run NONE -bench .
```