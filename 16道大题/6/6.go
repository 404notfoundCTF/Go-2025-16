package main // 声明主包

import (
	"fmt" // 导入格式化输出包
)

type Signature struct {
	message string
	nonce   string
}

// IdentifyInvalidSignatures 找出签名无效(签名发生重复)的消息，并返回对应的签名
func IdentifyInvalidSignatures(messages []string, signatures []Signature) []Signature {
	// 用于记录每条签名的消息数量
	messageCount := make(map[Signature]int)

	// 记录无效的签名
	var invalidMessages []Signature

	// 遍历消息和签名，统计每种签名的出现次数
	// ======== 请在此下方补充完整代码 ======
	nonceSeen := make(map[string]int)
	for i, sig := range signatures {
		if i < len(messages) {
			messageCount[sig]++
			nonceSeen[sig.nonce]++
			if nonceSeen[sig.nonce] > 1 {
				invalidMessages = append(invalidMessages, sig)
			}
		}
	}
	// ======== 请在此上方补充完整代码 ======

	return invalidMessages
}

func main() {
	// 示例代码
	messages := []string{"message1", "message2", "message3", "message4"}
	signatures := []Signature{
		{message: "message1", nonce: "nonce1"},
		{message: "message2", nonce: "nonce2"},
		{message: "message3", nonce: "nonce1"},
		{message: "message4", nonce: "nonce4"},
	}
	result := IdentifyInvalidSignatures(messages, signatures)
	for _, invalidSig := range result {
		fmt.Printf("%s %s\n", invalidSig.message, invalidSig.nonce)
	}
}
