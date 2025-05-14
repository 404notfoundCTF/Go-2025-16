package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateMerkleProof() (string, string) {
	txArray := []string{
		"d2EFCa744c6f2b567b1863dcF055C593afdC1178", //0
		"2905a6CA4Ebc3B5fc7541E115d67a890f9B00630", //1
		"3fC91A3afd70395Cd496C647d5a6CC9D4B2b7Ff8", //2
		"a23b31bc4aeaf5159ffc512a55549a3a39096b13", //3
	}
	//目标交易为txArray[0]，计算这个交易的merkle证明。返回值为两个16进制的string形式的哈希值（补全）
	hash2 := sha256.Sum256([]byte(txArray[1]))
	h2 := hex.EncodeToString([]byte(hash2[:]))

	hash3 := sha256.Sum256([]byte(txArray[2]))
	h3 := hex.EncodeToString([]byte(hash3[:]))
	hash4 := sha256.Sum256([]byte(txArray[3]))
	h4 := hex.EncodeToString([]byte(hash4[:]))
	hash34 := sha256.Sum256([]byte(h3 + h4))
	h34 := hex.EncodeToString([]byte(hash34[:]))
	fmt.Println("h2:", h2)
	fmt.Println("h34", h34)
	return h2, h34
}
func main() {
	generateMerkleProof()
}

//填写完毕
