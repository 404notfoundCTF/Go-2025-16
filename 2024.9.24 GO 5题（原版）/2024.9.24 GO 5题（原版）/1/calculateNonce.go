package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	//
	"strings" // 导入字符串处理包（引用后自动导入，无需手动）
	//
)

func CalculateNonce() int {
	data := "AK1r7dWdlxRiyyUfDyQB3Gb1KRCAAA3RyZaXtw"
	target := 3
	//在此处补充代码
	nonce := 30000                        // 初始 nonce 值，搜索从该值开始
	prefix := strings.Repeat("0", target) // 构造目标前导零字符串，例如 target=3 时为 "000"
	//Repeat 函数会把 0 这个字符串重复 target 次，生成一个新的字符串。
	//结束

	for {
		// 构造待哈希的数据，data 与 nonce 拼接
		hashData := fmt.Sprintf("%s%d", data, nonce)
		//fmt.Sprintf 就是用来“格式化字符串”，生成一个带变量内容的字符串，而不是直接打印出来。

		// 对数据计算 SHA-256 哈希
		hash := sha256.Sum256([]byte(hashData))

		// 将哈希字节转换为十六进制字符串（更易于比对）
		hexHash := hex.EncodeToString(hash[:])

		// 如果哈希值的前 target 位等于目标前导零字符串，则找到目标 nonce，退出循环
		if hexHash[:target] == prefix {

			break
		}
		// 否则继续尝试下一个 nonce，直到找到合适的为止
		nonce++
	}
	return nonce
}

// 执行
func main() {
	nonce := CalculateNonce()
	fmt.Println("找到的 nonce 是：", nonce) // 输出看到的 nonce 值
}
