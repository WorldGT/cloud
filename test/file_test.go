package test

import (
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片的大小
// const chunkSize = 100 * 1024 * 1024 // 100MB
const chunkSize = 1024 * 1024 // 1MB
// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("img/we.mp4")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile("img/we.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}
