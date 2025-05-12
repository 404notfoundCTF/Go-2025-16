/*
	根据一组交易地址，计算出对应的 Merkle 根。

1.每笔交易地址先单独做 SHA-256 哈希 → 得到“叶子节点”；
2.将相邻的哈希值拼接后再次哈希 → 得到“父节点”；
3.继续两两合并父节点并哈希 → 最终得到最顶层的哈希值，也就是 Merkle 根。
就像你把4个名字做成签名（哈希），两两签名叠一起再签名，
再把两个组合的签名合在一起签出最后的大签名——这个“大签名”就是你整组交易的数字指纹，
谁要验证其中任意一笔交易，只要部分签名就能搞定，无需全树。
*/
package main

import (
	"crypto/sha256" // 提供 SHA-256 哈希函数
	"encoding/hex"  // 用于将字节数组编码为十六进制字符串，便于输出和阅读
	"fmt"           // 用于格式化输出
)

// generateMerkle 函数用于计算给定交易数组的 Merkle 根
func generateMerkle() string {
	// 交易数组，每个元素代表一笔交易（这里用地址字符串简化模拟）
	txArray := []string{
		"388C818CA8B9251b393131C08a736A67ccB19297",
		"a6De4892df5410F44f73d85Cb941eb3D3c7d485a",
		"d81fa16C3C04106248bC3897A0fB4Eb5B21EDBbD",
		"8a2228705ec979961F0e16df311dEbcf097A2766",
	}

	// 第一步：对每笔交易做 SHA-256 哈希，得到叶节点
	hash1 := sha256.Sum256([]byte(txArray[0])) // 对第1笔交易做哈希
	h1 := hex.EncodeToString(hash1[:])         // 转为十六进制字符串
	hash2 := sha256.Sum256([]byte(txArray[1])) // 第2笔交易
	h2 := hex.EncodeToString(hash2[:])
	hash3 := sha256.Sum256([]byte(txArray[2])) // 第3笔交易
	h3 := hex.EncodeToString(hash3[:])
	hash4 := sha256.Sum256([]byte(txArray[3])) // 第4笔交易
	h4 := hex.EncodeToString(hash4[:])

	// 第二步：两两合并叶子节点哈希值，再做哈希，生成父节点
	hash12 := sha256.Sum256([]byte(h1 + h2)) // 合并 h1 和 h2，然后哈希
	h12 := hex.EncodeToString(hash12[:])     // 父节点 h12 的十六进制形式
	hash34 := sha256.Sum256([]byte(h3 + h4)) // 合并 h3 和 h4，然后哈希
	h34 := hex.EncodeToString(hash34[:])     // 父节点 h34 的十六进制形式

	// 第三步：继续向上合并，生成 Merkle 根
	hash1234 := sha256.Sum256([]byte(h12 + h34)) // 顶层节点，由 h12 和 h34 合并
	h1234 := hex.EncodeToString(hash1234[:])     // Merkle 根的十六进制表示

	// 打印最终的 Merkle 根
	fmt.Println(h1234, hex.EncodeToString(hash1234[:]))

	// 返回 Merkle 根（十六进制字符串）
	return h1234
}

// main 函数，程序入口
func main() {
	generateMerkle() // 调用函数生成并打印 Merkle 根
}
