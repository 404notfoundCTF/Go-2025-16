package main

import (
	"fmt"
)

func CountDuplicateSignatures(signature []string) int {
	countMap := make(map[string]int)

	//--------- 请在此下方补充代码---------
	for _, sig := range signature {
		countMap[sig]++
	}

	duplicateCount := 0
	for _, count := range countMap {
		if count > 1 {
			duplicateCount += count
		}
	}

	return duplicateCount
	//--------- 请在此上方补充代码---------
}

func main() {
	signatures := []string{"sig1", "sig2", "sig1", "sig3", "sig2"}
	result := CountDuplicateSignatures(signatures)
	fmt.Printf("重复签名消息的数量: %d\n", result)
}
