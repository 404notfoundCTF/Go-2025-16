package main

//
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateMerkle() string {
	txArray := []string{
		"388C818CA8B9251b393131C08a736A67ccB19297",
		"a6De4892df5410F44f73d85Cb941eb3D3c7d485a",
		"d81fa16C3C04106248bC3897A0fB4Eb5B21EDBbD",
		"8a2228705ec979961F0e16df311dEbcf097A2766",
	}
	//开始补全
	hash1 := sha256.Sum256([]byte(txArray[0]))
	h1 := hex.EncodeToString([]byte(hash1[:]))
	hash2 := sha256.Sum256([]byte(txArray[1]))
	h2 := hex.EncodeToString([]byte(hash2[:]))
	hash3 := sha256.Sum256([]byte(txArray[2]))
	h3 := hex.EncodeToString([]byte(hash3[:]))
	hash4 := sha256.Sum256([]byte(txArray[3]))
	h4 := hex.EncodeToString([]byte(hash4[:]))

	hash12 := sha256.Sum256([]byte(h1 + h2))
	h12 := hex.EncodeToString([]byte(hash12[:]))

	hash34 := sha256.Sum256([]byte(h3 + h4))
	h34 := hex.EncodeToString([]byte(hash34[:]))

	hash12354 := sha256.Sum256([]byte(h12 + h34))
	h1234 := hex.EncodeToString([]byte(hash12354[:]))

	fmt.Println(h1234)

	return h1234

	//结束
}

// 开始

func main() {
	generateMerkle()
}

//结束
