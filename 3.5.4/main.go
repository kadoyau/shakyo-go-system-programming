package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func textChunk(text string) io.Reader {
	byteText := []byte(text)
	crc := crc32.NewIEEE()
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteText)))
	// CRC計算とバッファへの書き込みを同時に行うMultiWriter
	writer := io.MultiWriter(&buffer, crc)
	// 2バイトめの5ビットめをを小文字にするとプライベート
	// teXtの2バイト目はe
	// ASCIIコードでeは01100101（2進数）
	// 5ビット目（左から数えて5番目）は1
	// ASCIIコードでは、大文字と小文字の違いは5ビット目で決まります
	// tEXtにすると2バイト目（E）の5ビット目は0
	io.WriteString(writer, "teXt")
	// 実際のデータをマルチライターに書き込む
	writer.Write(byteText)
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return &buffer
}
func dumpChunk(chunk io.Reader) {
	var length int32
	// file から 4バイト を読み取って
	// ビッグエンディアンとして解釈し
	// length（int32）に格納する
	// バイナリファイルのこのような部分：
	// 00 00 01 0A   ← 4バイトのデータ長（例: 266）
	// 49 48 44 52   ← チャンクタイプ（例: "IHDR"）
	// があったとき、binary.Read の後には length = 266 になり、
	// ファイルポインタは “49”（チャンクタイプ）に位置している状態になります。
	// [4B: length] [4B: type] [N B: data] [4B: CRC]
	//  ↑              ↑
	//  |              └─ chunk.Read(buffer) の位置
	//  └─ binary.Read(..., &length) が読み取った
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
	if bytes.Equal(buffer, []byte("teXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}
func readChunks(file *os.File) []io.Reader {
	// チャンクを格納する配列
	var chunks []io.Reader
	// 最初の8 バイトを飛ばす
	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks,
			io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に移動
		// 現在位置は長さを読み終わった箇所なので
		// チャンク名(4 バイト) + データ長 + CRC(4 バイト) 先に移動
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}
func main() {
	file, err := os.Open("PNG_transparency_demonstration_1.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("PNG_transparency_demonstration_secret.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := readChunks(file)
	// シグニチャ書き込み
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// 先頭に必要なIHDR チャンクを書き込み
	io.Copy(newFile, chunks[0])
	// テキストチャンクを追加
	io.Copy(newFile, textChunk("Lambda Note++"))
	// 残りのチャンクを追加
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	file2, err2 := os.Open("PNG_transparency_demonstration_secret.png")
	if err2 != nil {
		panic(err2)
	}
	defer file2.Close()
	chunks2 := readChunks(file2)
	for _, chunk := range chunks2 {
		dumpChunk(chunk)
	}
}
