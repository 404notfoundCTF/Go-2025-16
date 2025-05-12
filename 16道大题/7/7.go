package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateMessageHash(message string) string {
	// 将输入字符串转换为字节序列
	data := []byte(message)

	// 使用 SHA256 算法计算哈希值

	//========== 请在下方补充完整代码==========
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])

	return hashStr[:5]
	//========== 请在上方补充完整代码==========
}

func main() {
	message := "Hello, Blockchain!"
	hash := CalculateMessageHash(message)
	fmt.Printf("Message: %s\nHash (first 5 chars): %s\n", message, hash)
}
