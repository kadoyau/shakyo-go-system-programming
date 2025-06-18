package main

import (
	"io"
	"os"
)

// 古いファイル（old.txt）を新しいファイル（new.txt）にコピーしてみましょう。
// 本章で紹介したサンプルコードを応用すれば難しくないと思います。
// さらに改造して実用的なコマンドにしてみたいと思われる方は、コマンドラインオ
// プションでファイル名を渡せるようにするとよいでしょう。本書の範囲からは外れ
// るので詳細は省きますが、os.Args という文字列配列にオプションが格納されます。
// また、標準ライブラリにあるflag パッケージを使うと、オプションのパース処理が
// より便利に行えます。
func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, oldFile)
}
