/* 以下代码模拟了一个小型的区块链交易列表，然后：

对第2笔交易（tx2）进行哈希，这是我们要证明的目标交易；

对第3、4笔交易分别哈希后合并，再哈希一次，形成一个叫 hash34 的中间节点；

最后，它把 tx2 的哈希 和 hash34 打印出来并返回。

这两个哈希值就是所谓的 Merkle 证明，你可以用它们去验证 tx2 是真的在这棵交易树里，而不需要整棵树的数据。 */

package main

import (
	"crypto/sha256" // 用于哈希计算
	"encoding/hex"  // 用于将哈希值转换为十六进制字符串
	"fmt"           // 用于输出调试信息
)

// generateMerkleProof 函数生成并返回某一笔交易（此处为tx2）在 Merkle 树中的证明
func generateMerkleProof() (string, string) {
	// 定义交易数组，每个元素代表一笔交易（简化为地址字符串）
	txArray := []string{
		"d2EFCa744c6f2b567b1863dcF055C593afdC1178", // tx1
		"2905a6CA4Ebc3B5fc7541E115d67a890f9B00630", // tx2
		"3fC91A3afd70395Cd496C647d5a6CC9D4B2b7Ff8", // tx3
		"a23b31bc4aeaf5159ffc512a55549a3a39096b13", // tx4
	}

	// 对 tx2 进行哈希计算，得到叶子节点的哈希值 hash2
	hash2 := sha256.Sum256([]byte(txArray[1]))
	h2 := hex.EncodeToString(hash2[:]) // 将 hash2 转为十六进制字符串

	// 对 tx3 进行哈希计算，得到 hash3
	hash3 := sha256.Sum256([]byte(txArray[2]))
	h3 := hex.EncodeToString(hash3[:]) // 转换为字符串形式

	// 对 tx4 进行哈希计算，得到 hash4
	hash4 := sha256.Sum256([]byte(txArray[3]))
	h4 := hex.EncodeToString(hash4[:]) // 转换为字符串形式

	// 合并 hash3 和 hash4，生成它们的父节点 hash34
	hash34 := sha256.Sum256([]byte(h3 + h4))
	h34 := hex.EncodeToString(hash34[:]) // 转为字符串

	// 输出中间结果，方便调试
	fmt.Println(h2, h34)

	// 返回 Merkle 证明所需的两个兄弟节点哈希值：hash2（tx2的哈希）和 hash34（tx3、tx4的合并哈希）
	return h2, h34
}

func main() {
	generateMerkleProof() // 调用函数生成并打印 Merkle 证明
}
