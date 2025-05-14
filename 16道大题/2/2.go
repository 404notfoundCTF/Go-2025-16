package main

import "fmt"

func CountValidPBFTSystems(matrix [][2]int) int {
	count := 0
	//请在下方补全代码=============
	for _, system := range matrix {
		n := system[0]     // 总节点数
		fMax := system[1]  // 拜占庭节点数
		if 3*fMax+1 <= n { // 检查 PBFT 条件
			count++
		}
	}
	//请在上方补全代码=============
	return count
}

func main() {
	testMatrix := [][2]int{
		{4, 1},
		{7, 2},
		{3, 0},
	}
	result := CountValidPBFTSystems(testMatrix)
	fmt.Printf("满足 PBFT 条件的系统数量: %d\n", result)
}
