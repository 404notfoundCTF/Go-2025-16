package main

import "fmt"

// CountNonEmptySlots 计算一个分支节点中非空子节点的数目
// 输入为一个十进制整数，该整数可以展开为16个二进制位，每个值为1表示对应槽位非空。
func CountNonEmptySlots(num int) int {
	count := 0

	//============请在下方补充代码============
	for i := 0; i < 16; i++ {
		// 检查每一位是否为1，使用位运算
		if (num & (1 << i)) != 0 {
			count++
		}
	}

	//===========请在上方补充代码==============
	return count
}

func main() {
	// 测试用例1：num = 7，二进制为0000000000000111，预期输出3
	fmt.Println("Test case 1:", CountNonEmptySlots(7)) // 预期输出: 3

	// 测试用例2：num = 0，二进制为0000000000000000，预期输出0
	fmt.Println("Test case 2:", CountNonEmptySlots(0)) // 预期输出: 0

	// 测试用例3：num = 65535，二进制为1111111111111111，预期输出16
	fmt.Println("Test case 3:", CountNonEmptySlots(65535)) // 预期输出: 16
}
