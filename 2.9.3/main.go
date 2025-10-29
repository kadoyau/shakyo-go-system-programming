package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request:", r.Method, r.URL)
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "world",
	}
	// JSONの文字列変換、gzip圧縮を行いながら、圧縮前の出力を標準出力にも出す

	// gzip.Writerを作成し、http.ResponseWriterにラップする
	gz := gzip.NewWriter(w)
	// 圧縮処理はgzip.Writerでバッファリングされるので、closeを読んでデータを出力した上でリソースを解放する
	defer gz.Close()

	// io.MultiWriterを使ってgzip.Writerと標準出力に同時に出力
	multiWriter := io.MultiWriter(gz, os.Stdout)
	// JSON encoderでsourceをJSONにしてmultiWriterに書き込む
	err := json.NewEncoder(multiWriter).Encode(source)
	if err != nil {
		// エンコードに失敗した場合、エラーメッセージを返す
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

}

func main() {
	log.Println("Starting server on :8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
