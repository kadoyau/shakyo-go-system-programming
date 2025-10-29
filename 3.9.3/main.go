package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

// archive/zip パッケージを使ってzip ファイルを作成してみましょう。出力先の
// ファイルのWriter （以下のコードのfile）をまず作って、それをzip.NewWriter()
// 関数に渡すと、zip ファイルの書き込み用の構造体ができます（図3.9「zip ファイル
// の読み書きに構造体そのものではなくインタフェースを使う」）。最後にClose() を
// 確実に呼ぶ必要がありますが、これにはGo 言語のdefer という機能を使って次のよ
// うにすればいいでしょう。
// zipWriter := zip.NewWriter(file)
// defer zipWriter.Close()
// この構造体そのものはio.Writer ではありませんが、Create() メソッドを呼ぶ
// と、個別のファイルを書き込むためのio.Writer が返ってきます。
// writer, err := zipWriter.Create("newfile.txt")
// 上記の例では、newfile.txt という実際のファイルが、最初に作った出力先の
// ファイルfile へと圧縮されます。では、実際のファイルではなく、文字列strings.
// Reader を使ってzip ファイルを作成するにはどうすればいいでしょうか。考えてみ
// てください。
func main() {
	outputFile, err := os.Create("393_out.zip")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	zipWriter := zip.NewWriter(outputFile)
	defer zipWriter.Close()

	// zip中のnewfile.txtにHello, worldと書き込む
	sReader := strings.NewReader("Hello, world!")
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(writer, sReader)
	if err != nil {
		panic(err)
	}

}
