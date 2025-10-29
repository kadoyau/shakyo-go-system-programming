package main

import (
	"crypto/rand"
	"io"
	"os"
)

// ファイルを作成してランダムな内容で埋めてみましょう。
// crypto/rand パッケージ（本来は付録A で紹介するように暗号用の機能）をイン
// ポートすると、rand.Reader というio.Reader が使えます。このReader は、ラン
// ダムなバイトを延々と出力し続ける無限長のファイルのような動作をします。これを
// 使って、1024 バイトの長さのバイナリファイルを作ってみましょう。
// ヒントですが、io.Copy() を使ってはいけません。io.Copy() はReader の終了
// まですべて愚直にコピーしようとします。
func main() {
	outputFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	// 1024 バイトの長さのバイナリファイルを作成
	if _, err := io.CopyN(outputFile, rand.Reader, 1024); err != nil {
		panic(err)
	}

	// ファイルを閉じる
	defer outputFile.Close()
}
