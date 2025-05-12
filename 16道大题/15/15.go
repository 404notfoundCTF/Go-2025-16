package main

import (
	"fmt"
	"strings"
)

// 找出满足target个前缀零个数的哈希值，返回判断结果的布尔变量
func SortPoWHash(hash string, target int) bool {
	//=============请在下方补全代码=============
	prefix := 0
	for _, char := range hash {
		if char != '0' {
			break
		}
		prefix++

	}
	return prefix >= target
	//=============请在上方补全代码=============
}

func main() {
	// 示例哈希值
	hash1 := "0000f8a7b91e11e41f25e85fac067825b0066c8352e3fe2978f8d765ac8792bd"
	hash2 := "00f8a7b91e11e41f25e85fac067825b0066c8352e3fe2978f8d765ac8792bd56"
	hash3 := "f8a7b91e11e41f25e85fac067825b0066c8352e3fe2978f8d765ac8792bd56ab"

	// 测试不同的目标值
	fmt.Println("哈希值:", hash1)
	fmt.Println("目标值为3时:", SortPoWHash(hash1, 3)) // 应该返回 true (4个0 >= 3)
	fmt.Println("目标值为4时:", SortPoWHash(hash1, 4)) // 应该返回 true (4个0 >= 4)
	fmt.Println("目标值为5时:", SortPoWHash(hash1, 5)) // 应该返回 false (4个0 < 5)

	fmt.Println("\n哈希值:", hash2)
	fmt.Println("目标值为1时:", SortPoWHash(hash2, 1)) // 应该返回 true (2个0 >= 1)
	fmt.Println("目标值为2时:", SortPoWHash(hash2, 2)) // 应该返回 true (2个0 >= 2)
	fmt.Println("目标值为3时:", SortPoWHash(hash2, 3)) // 应该返回 false (2个0 < 3)

	fmt.Println("\n哈希值:", hash3)
	fmt.Println("目标值为1时:", SortPoWHash(hash3, 1)) // 应该返回 false (0个0 < 1)

	// 演示实际区块链难度案例
	fmt.Println("\n区块链实际案例:")
	difficulty := 4
	validHash := strings.Repeat("0", difficulty) + "19fde3955bd3ce52b9073d1c9581b988f4cea9ae2c2f8a8c3c9afb706e86d7a4"
	invalidHash := strings.Repeat("0", difficulty-1) + "a9fde3955bd3ce52b9073d1c9581b988f4cea9ae2c2f8a8c3c9afb706e86d7a4"

	fmt.Printf("难度目标为%d时，哈希值为%s的结果: %v\n", difficulty, validHash, SortPoWHash(validHash, difficulty))
	fmt.Printf("难度目标为%d时，哈希值为%s的结果: %v\n", difficulty, invalidHash, SortPoWHash(invalidHash, difficulty))
}
