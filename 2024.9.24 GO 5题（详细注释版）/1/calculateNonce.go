/* 这个程序大致就是在模拟区块链中“挖矿”的过程，用 Go 语言写的一个简单版“工作量证明（Proof of Work）”。
想象你要解决一个谜题，规则是这样的：

给你一个字符串（叫 data），你要往后面加一个数字（叫 nonce）。

把这个“字符串+数字”拼起来，然后用一个特殊的算法（SHA-256）加密。

如果加密后的结果（哈希值）前面刚好有指定数量的“0”，那你就赢了！

否则你就换个数字，再试一次，直到猜中为止。

这就像在玩一个猜数字的游戏，但你只能一个个穷尽试过去。 */

package main //默认格式，导入包名

import (
	"crypto/sha256" // 导入加密包，用于计算 SHA-256 哈希值
	"encoding/hex"  // 导入编码包，用于将哈希字节转换为十六进制字符串
	"fmt"           // 导入格式化输出包

	//℘℘℘ 在此间自动导入↓（无提示）
	"strings" // 导入字符串处理包（引用后自动导入，无需手动）
	//℘℘℘ 在此间自动导入↑
)

// CalculateNonce 用于通过简单的工作量证明机制（类似区块链挖矿）来计算一个满足条件的 nonce 值
func CalculateNonce() int {
	data := "AK1r7dWdlxRiyyUfDyQB3Gb1KRCAAA3RyZaXtw" // 参与计算哈希的数据内容
	target := 3                                      // 目标哈希前导零的数量（表示难度）

	//℘℘℘ 在此间补充代码↓（不会有任何提示，联系上下文）
	nonce := 30000                        // 初始 nonce 值，搜索从该值开始
	prefix := strings.Repeat("0", target) // 构造目标前导零字符串，例如 target=3 时为 "000"
	//Repeat 函数会把 0 这个字符串重复 target 次，生成一个新的字符串。
	//℘℘℘ 在此间补充代码↑

	for {
		// 构造待哈希的数据，data 与 nonce 拼接
		hashData := fmt.Sprintf("%s%d", data, nonce)
		//fmt.Sprintf 就是用来“格式化字符串”，生成一个带变量内容的字符串，而不是直接打印出来。

		// (手动测试，非原题语句)打印调试信息：原始字符串
		//fmt.Println(hashData)

		// （手动测试，非原题语句）打印调试信息：转换为字节数组
		//fmt.Println([]byte(hashData))

		// 对数据计算 SHA-256 哈希
		hash := sha256.Sum256([]byte(hashData))

		// （手动测试，非原题语句）打印调试信息：输出原始哈希值（[32]byte 类型）
		//fmt.Println(hash)

		// 将哈希字节转换为十六进制字符串（更易于比对）
		hexHash := hex.EncodeToString(hash[:])

		// 如果哈希值的前 target 位等于目标前导零字符串，则找到目标 nonce，退出循环
		if hexHash[:target] == prefix {
			break
		}
		// 否则继续尝试下一个 nonce，直到找到合适的为止
		nonce++
	}

	// 返回找到的 nonce 值
	return nonce

}
func main() {
	nonce := CalculateNonce()          // 调用POW算法计算nonce值函数，赋值给变量 nonce （找到的 nonce 值）
	fmt.Println("找到的 nonce 是：", nonce) // 输出看到的 nonce 值
}

/*
程序的作用：
1.从 30000 开始尝试一个数字 nonce；

2.把这个数字和字符串拼接在一起；

3.对拼接后的内容进行 SHA-256 哈希加密；

4.查看加密后的结果是不是以目标数量的“0”开头；

5.如果不是，继续尝试下一个数字；

6.如果是，程序返回这个数字（nonce），说明你“挖到矿”了！ */
