package main

import (
	"fmt"
	"strings"
)

// ValidateAddressLength 验证以太坊地址是否为 20 字节（40 个十六进制字符），返回true或false
func ValidateAddressLength(address string) bool {
	//===================请在下方补全代码=====================
	if strings.HasPrefix(address, "0x") || strings.HasPrefix(address, "0X") {
		address = address[2:]
	}
	return len(address) == 40
	//===================请在上方补全代码=====================
}

func main() {
	// 验证示例
	addresses := []string{
		"0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",   // 有效地址（带前缀）
		"5B38Da6a701c568545dCfcB03FcB875f56beddC4",     // 有效地址（无前缀）
		"0x5B38Da6a701c568545dCfcB03FcB875f56beddC",    // 无效地址（长度不足）
		"0x5B38Da6a701c568545dCfcB03FcB875f56beddC4AB", // 无效地址（长度过长）
	}

	for _, addr := range addresses {
		result := ValidateAddressLength(addr)
		fmt.Printf("地址: %s\n结果: %v\n\n", addr, result)
	}
}
