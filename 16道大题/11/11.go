package main

import (
	"fmt"
)

// 将目标字符转换为rune类型
func CountCharacterInAddress(address string, targetChar string) int {
	target := []rune(targetChar)[0]
	count := 0

	//==========请在此处补充代码==========
	for _, char := range address {
		if char == target {
			count++
		}
	}
	//==========请在此处补充代码==========

	return count
}

func main() {
	// 测试用例
	address := "1A2b3C4d5E6f7G8h9I0j"
	targetChar := "A"
	result := CountCharacterInAddress(address, targetChar)
	fmt.Printf("地址 %s 中字符 %s 出现的次数: %d\n", address, targetChar, result)
}
