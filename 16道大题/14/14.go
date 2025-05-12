package main

import (
	"fmt"
)

// 区块结构体定义
type Block struct {
	BlockName string
	Hash      string
	PrevHash  string
}

// =============以下为题目源码及注释=====================

// 根据区块存储的哈希值进行区块排序，返回排序数组（以区块名为元素）
func SortBlocks(blocks []Block) []string {
	// 找到起始区块（设定为存储的前一个哈希为空）
	var startBlock Block
	for _, block := range blocks {
		if block.PrevHash == "null" {
			startBlock = block
			break
		}
	}
	//=============以上为题目源码及注释=====================

	//=============以下为题目源码及注释=====================
	// 构建区块链顺序
	var orderedBlocks []string
	currentBlock := startBlock
	//=============以上为题目源码及注释=====================

	//=============请在下方补全代码================

	// 将起始区块加入有序数组
	orderedBlocks = append(orderedBlocks, currentBlock.BlockName)

	// 创建哈希映射，便于快速查找
	blockMap := make(map[string]Block)
	for _, block := range blocks {
		blockMap[block.Hash] = block
	}

	// 通过前一个区块的哈希值连接下一个区块
	for len(orderedBlocks) < len(blocks) {
		// 遍历所有区块，查找当前区块哈希值作为PrevHash的下一个区块
		found := false
		for _, nextBlock := range blocks {
			if nextBlock.PrevHash == currentBlock.Hash {
				// 找到了下一个区块
				orderedBlocks = append(orderedBlocks, nextBlock.BlockName)
				currentBlock = nextBlock
				found = true
				break
			}
		}

		if !found {
			break // 如果没有找到下一个区块，就退出循环
		}
	}

	return orderedBlocks
	//=============请在上方补全代码================
}

func main() {
	// 创建测试用的区块数据
	blocks := []Block{
		{BlockName: "Block3", Hash: "hash3", PrevHash: "hash2"},
		{BlockName: "Block1", Hash: "hash1", PrevHash: "null"},
		{BlockName: "Block2", Hash: "hash2", PrevHash: "hash1"},
		{BlockName: "Block4", Hash: "hash4", PrevHash: "hash3"},
	}

	// 排序区块并输出结果
	orderedBlocks := SortBlocks(blocks)
	fmt.Println("排序后的区块链顺序:", orderedBlocks)
}
