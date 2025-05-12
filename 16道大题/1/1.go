package main

import "fmt"

func IsVotePassed(number int) bool {
	//============请在下方补充代码==============
	if number < 0 || number > 31 {
		return false
	}
	count := 0
	for num := number; num > 0; num >>= 1 {
		if num&1 == 1 {
			count++
		}

	}
	return count >= 3
	//============请在上方补充代码==============
}

func main() {
	// 测试用例
	fmt.Println(IsVotePassed(21)) // 二进制 10101，3个1，true
	fmt.Println(IsVotePassed(3))  // 二进制 00011，2个1，false
	fmt.Println(IsVotePassed(31)) // 二进制 11111，5个1，true
	fmt.Println(IsVotePassed(32)) // 超出5位，false
}
