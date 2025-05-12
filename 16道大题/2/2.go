package main

import "fmt"

func CountValidPBFTSystems(matrix [][2]int) int {
	count := 0
	//请在下方补全代码=============
	for i := 0; i < len(matrix); i++ {
		n := matrix[i][0]
		fMax := matrix[i][1]
		for f := 0; f <= fMax; f++ {
			if 3*f+1 <= n {
				count++
			}
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
