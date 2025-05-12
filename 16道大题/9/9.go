package main

import "fmt"

// 提取哈希值的前8个字符作为地址前缀
// 返回地址前缀，形如"abcdef12"
func GenerateAddressPrefix(hashValue string) string {
	//--------- 请在此下方补充代码---------
	if len(hashValue) < 8 {
		return hashValue
	}
	return hashValue[:8]
	//--------- 请在此上方补充代码---------
}

func main() {
	hashValue := "abcdef1234567890"
	prefix := GenerateAddressPrefix(hashValue)
	fmt.Println("哈希值:", hashValue)
	fmt.Println("地址前缀:", prefix)
}
