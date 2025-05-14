package main

import "fmt"

func CountValidPBFTSystems(matrix [3][2]int) int {
	count := 0
	//=============请在下方补全代码=============
	for _, system := range matrix {
		n := system[0]     // 总节点数
		fMax := system[1]  // 拜占庭节点数
		if 3*fMax+1 <= n { // 通过公式检查 PBFT 条件（推导过程看md文档）,基本公式 3*fMax + 1 <= n 来源于 PBFT 对一致性的数学需求，确保正常节点数量至少占总节点的 2/3 以上，并留有余量以应对恶意行为。这个条件的推导基于分布式系统的容错理论，是 PBFT 算法设计的核心。
			count++
		}
	}
	//=============请在上方补全代码=============
	return count
}

func main() {
	testMatrix := [3][2]int{
		{4, 1}, //{4, 1}：3*1 + 1 = 4 ≤ 4，满足条件。
		{7, 2}, //{7, 2}：3*2 + 1 = 7 ≤ 7，满足条件。
		{3, 0}, //{3, 0}：3*0 + 1 = 1 ≤ 3，满足条件。
	}
	result := CountValidPBFTSystems(testMatrix)
	fmt.Printf("%d\n", result) //预计输出3
}
