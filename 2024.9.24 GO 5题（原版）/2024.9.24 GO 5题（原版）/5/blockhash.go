package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"

	//开始补全
	"fmt"
	//结束补全
)

type BlockHeader struct {
	Version       uint64
	PrevBlockHash []byte
	MerkleRoot    []byte
	Timestamp     time.Time
	Nonce         uint64
}

const (
	fixedTimestamp int64 = 1614363477
)

func CalculateBlockHash() [32]byte {
	header := BlockHeader{
		Version:       1,
		PrevBlockHash: []byte{111, 205, 99, 52, 176, 155, 165, 121, 152, 53, 180, 205, 245, 121, 207, 182, 147, 128, 55, 178, 224, 247, 64, 123, 3, 79, 198, 60, 73, 145, 68, 87},
		MerkleRoot:    []byte{0, 5, 124, 54, 190, 235, 71, 37, 145, 71, 158, 247, 39, 193, 43, 216, 235, 145, 122, 97, 161, 184, 169, 152, 126, 149, 45, 191, 248, 59, 251, 149},
		Timestamp:     time.Unix(fixedTimestamp, 0),
		Nonce:         67890,
	}
	buf := new(bytes.Buffer)
	//在此处补充代码
	//Version需要用小端形式放入buf中
	binary.Write(buf, binary.LittleEndian, header.Version)
	//完毕
	buf.Write(header.PrevBlockHash)
	buf.Write(header.MerkleRoot)
	binary.Write(buf, binary.LittleEndian, header.Timestamp.Unix())
	binary.Write(buf, binary.LittleEndian, header.Nonce)
	hash := sha256.Sum256(buf.Bytes())
	fmt.Println("buf:", buf.Bytes())

	return hash
}

// 调用
func main() {
	fmt.Print("hash:", CalculateBlockHash())
}
