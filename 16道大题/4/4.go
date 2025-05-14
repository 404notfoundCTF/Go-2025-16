package main

import (
	"fmt"
	"sort"
)

func SortPoWMiners(a, b, c, d float64) (float64, float64, float64, float64) {
	// 将输入构造成一个切片
	hashPowers := []float64{a, b, c, d}

	//=========请在下方补充完整代码=========
	// 对切片按照算力从高到低排序
	sort.Slice(hashPowers, func(i, j int) bool {
		return hashPowers[i] > hashPowers[j]
	})

	// 返回排序后的结果
	return hashPowers[0], hashPowers[1], hashPowers[2], hashPowers[3]
	//=========请在上方补充完整代码==========
}

func main() {
	// 测试用例
	a, b, c, d := SortPoWMiners(15.0, 35.0, 5.0, 25.0)
	fmt.Printf("%.1f %.1f %.1f %.1f\n", a, b, c, d)
}
