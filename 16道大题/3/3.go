package main

import (
	"fmt"
	"sort" // 导入排序包，用于排序操作
)

type Node struct { // 定义 Node 结构体
	Name        string // 节点名称
	StakingAge  int    // 权益年龄
	TokenAmount int    // 通证数量
}

func SortPOSNodesTest(node [3]*Node) (string, string, string) {
	// 将数组转换为切片以便排序
	nodes := node[:]

	// 自定义排序规则，按权益值降序排序
	sort.Slice(nodes, func(i, j int) bool {
		//===========请在下方补充代码==========
		// 计算节点权力值
		valueI := nodes[i].StakingAge * nodes[i].TokenAmount
		valueJ := nodes[j].StakingAge * nodes[j].TokenAmount
		if valueI != valueJ {
			return valueI > valueJ
		}

		// 权益值相同时的次级排序（按名称升序）
		return nodes[i].Name < nodes[j].Name
		//===========请在上方补充代码==========
	})

	// 返回排序节点的节点名称
	return nodes[0].Name, nodes[1].Name, nodes[2].Name
}

func main() { // 示例代码
	nodes := [3]*Node{
		&Node{Name: "NodeA", StakingAge: 10, TokenAmount: 50}, //NodeA: 10 * 50 = 500
		&Node{Name: "NodeB", StakingAge: 20, TokenAmount: 30}, //NodeB: 20 * 30 = 600
		&Node{Name: "NodeC", StakingAge: 15, TokenAmount: 40}, //NodeC: 15 * 40 = 600
	}
	name1, name2, name3 := SortPOSNodesTest(nodes)
	fmt.Printf("%s %s %s\n", name1, name2, name3) //NodeB (600), NodeC (600), NodeA (500)，其中 NodeB 和 NodeC 按名称升序
}
