<div align="right">
    <a href="./README.md">English</a> | <a href="./README_jp.md">日本語</a>
</div>

# gomarkdown
Go言語でマークダウンのテキストをHTMLの文字列に置き換える関数です。
とても遅いです。

## インストール
以下のコマンドでダウンロードします。
```
go get -u github.com/TomSuzuki/gomarkdown
```
その後以下のようにパッケージのインポートを書きます。
```
import "github.com/TomSuzuki/gomarkdown"
```
バージョンを指定したい場合は`go get`時に、一番後ろに`@vX.X.X`のような指定を加えるとできたはずです。

## 使用方法
```
md, _ = ioutil.ReadFile("./markdown.md")
html := gomarkdown.MarkdownToHTML(string(md))
```
細かい設定機能は存在しません。後で作成するかもです。

## ライセンス
- [MIT License](https://ja.wikipedia.org/wiki/MIT_License)
