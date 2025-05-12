/* 这个程序定义并构造了一个区块头（BlockHeader），包括版本号、前一区块哈希、Merkle 根、时间戳和 Nonce。
然后它将这些字段编码为一个连续的字节流（按照区块链协议格式），
并对整个字节流执行 SHA-256 哈希计算，最终生成该区块头的唯一标识 —— 区块哈希。
这个哈希值是区块链中用来识别并链接每个区块的核心机制。 */

package main

import (
	"bytes"           // 提供字节缓冲功能，用于拼接数据
	"crypto/sha256"   // 提供 SHA-256 哈希函数
	"encoding/binary" // 用于将数据以特定字节序编码成字节序列
	"fmt"             // 用于格式化输出
	"time"            // 提供时间处理
)

// BlockHeader 结构体表示一个区块头
type BlockHeader struct {
	Version       uint64    // 区块版本号
	PrevBlockHash []byte    // 前一个区块的哈希
	MerkleRoot    []byte    // 本区块的 Merkle 根哈希
	Timestamp     time.Time // 区块生成时间
	Nonce         uint64    // 工作量证明用的随机数
}

const (
	fixedTimestamp int64 = 1614363477 // 固定时间戳（模拟数据用）
)

// CalculateBlockHash 用于计算一个模拟区块头的哈希值
func CalculateBlockHash() [32]byte {
	// 构造一个区块头（模拟数据）
	header := BlockHeader{
		Version:       1,                                                                                                                                                          // 设置区块版本为 1
		PrevBlockHash: []byte{111, 205, 99, 52, 176, 155, 165, 121, 152, 53, 180, 205, 245, 121, 207, 182, 147, 128, 55, 178, 224, 247, 64, 123, 3, 79, 198, 60, 73, 145, 68, 87}, // 前一区块哈希
		MerkleRoot:    []byte{0, 5, 124, 54, 190, 235, 71, 37, 145, 71, 158, 247, 39, 193, 43, 216, 235, 145, 122, 97, 161, 184, 169, 152, 126, 149, 45, 191, 248, 59, 251, 149},  // Merkle 根
		Timestamp:     time.Unix(fixedTimestamp, 0),                                                                                                                               // 使用固定的 Unix 时间戳
		Nonce:         67890,                                                                                                                                                      // 指定 Nonce
	}

	buf := new(bytes.Buffer) // 创建一个新的字节缓冲区，用于拼接区块头字段
	//开始补全
	binary.Write(buf, binary.LittleEndian, header.Version) // 按小端序写入版本号（version需要用小端形式放入buf中）
	//结束补全

	buf.Write(header.PrevBlockHash) // 写入前一个区块的哈希值
	buf.Write(header.MerkleRoot)    // 写入 Merkle 根哈希值

	binary.Write(buf, binary.LittleEndian, header.Timestamp.Unix()) // 写入时间戳（小端）
	binary.Write(buf, binary.LittleEndian, header.Nonce)            // 写入 Nonce（小端）

	hash := sha256.Sum256(buf.Bytes()) // 对整个区块头字节序列做 SHA-256 哈希

	fmt.Println("buf:", buf.Bytes()) // 打印拼接后的字节数据（调试用）

	return hash // 返回计算出的哈希值
}

func main() {
	fmt.Print("hash:", CalculateBlockHash()) // 调用函数并打印结果
}

/* 说人话就是：

这个程序像是给一个“区块身份证”做指纹识别：

它准备了一份身份证材料（版本号、家族关系、交易信息、出生时间、验证码）；

把这些内容拼在一起做了一次“指纹采集”（SHA-256）；

然后打印出这个区块的“指纹”——也就是哈希值。

这个哈希值可以被用来唯一地代表这个区块，就像区块链里每一个区块都有一个独一无二的身份证号一样，

谁也造不了假，因为只要里面有一丁点改动，这个哈希就会完全不同。 */
