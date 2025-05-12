package main

import "fmt"

// 计算并返回交易费率数组
func TransactionFeeRate(totalFees, transactionSizes []int) []int {
	// 初始化一个存放手续费率的数组
	feeRates := make([]int, len(totalFees))

	// 遍历计算每个交易的手续费率
	//===========请在下方补充代码=============
	for i := 0; i < len(totalFees); i++ {
		if transactionSizes[i] != 0 { // 防止除以零
			feeRates[i] = totalFees[i] * 100 / transactionSizes[i] //手续费率 = (总费用 * 100) / 交易金额
		}
	}
	//===========请在上方补充代码=============

	// 最后计算平均手续费率
	return feeRates
}

func main() {
	totalFees := []int{100, 200, 300}
	transactionSizes := []int{1000, 2000, 1500}
	result := TransactionFeeRate(totalFees, transactionSizes)
	fmt.Println(result) // 应该输出 [10 10 20]
}
