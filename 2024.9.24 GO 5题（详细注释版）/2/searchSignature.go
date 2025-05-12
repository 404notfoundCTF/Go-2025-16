//以下代码"转换"，“哈希”，“签名”过程要求全部掌握，大范围逻辑填空

/* 本程序使用以太坊加密库，对两条消息分别使用指定私钥进行签名。
签名后生成的签名值会以十六进制格式与已有签名列表进行比对，
查找是否存在匹配项。整个流程展示了如何完成消息签名、签名编码，以及签名验证的基本操作，是实现交易签名、消息认证的基础。 */

package main // 定义 main 包，表示可执行程序入口

import ( // 导入依赖包
	"encoding/hex" // 用于进行16进制编码和解码 // 用于进行16进制编码和解码
	"fmt"          // 用于格式化输出 // 用于输出格式化的文本
	"log"          // 用于日志记录，错误时输出并终止程序 // 用于记录错误信息

	"github.com/ethereum/go-ethereum/crypto" // 导入 go-ethereum 中的 crypto 包，用于加密、签名以及哈希计算 // 导入用于ECDSA签名与Keccak256哈希的加密库
)

func searchSignature() (string, string) { // 定义 searchSignature 函数，返回两个字符串结果 // 定义一个函数，返回匹配到的两个签名的十六进制字符串
	privateKeyHex := "6e0aac9f7bf63010390545b055c23505259a2ce4c87ddbf3acf48cefd83a6dac" // 定义私钥的16进制字符串 // 定义待使用的私钥16进制字符串
	messageOne := "hello world"                                                         // 定义消息1 // 第一个要签名的消息
	messageTwo := "hello blockchain"                                                    // 定义消息2 // 第二个要签名的消息
	sigArray := []string{                                                               // 定义预先存储的签名数组，用于匹配验证 // 用于保存若干预期签名的字符串集合
		"3cc4af10427f9b0cdb196fa3542ccc1f90ae11f5296ac666bea0f44f935480ff7ab212d01a1c1e5245a7a099f4b64e2de47008f8626d4fcb172facc5e087dd3b00", // 签名1 // 签名样本1
		"d633b72ee1f04e70a7d9043e6c8bcc8b64920ffb24a8dcd27ce487741c94442b2cc3899d553734550e51eb6b6755861578397b892d5308497d4f7c787b9752a700", // 签名2 // 签名样本2
		"458be524741b15bd3c665f1ef5d9564ec95935adf0957606db11759bf5e162ab5b53f63d490e91292c61250f6cf4ee96460090a79ecc21f95e5bb7b04a2e793900", // 签名3 // 签名样本3
		"aa17c15fd8618ec2623c776214d6e2606d7f80574b5af026fe8c299fad52364668b761293e70cdc28530e878cc8e94e12981f24de98853c58f32929c4864c5c801", // 签名4 // 签名样本4
		"bf8138c93e565985c4d48dbe00ddc45c12c30ecc862290cb009ea452b3faf589325a2454c739cda4180d3c7bc1f63ec7e529636a4b1815d2c252c93e14281cdf01", // 签名5 // 签名样本5
		"a73ad2c912b1d59827bccd70042f74b403eaad6c1dcad2e4d2a01166eb7e6d5c5a0669739e7380270d888627e5074453413a6b224d038e5d10717f56555c71d900", // 签名6 // 签名样本6
		"75c362d664e8f79a1cdf4524b5ae8dec1a4541500b5ab5738e093000e160e79025cb103b56f5d6d28bac627d0cee0507b3c0f1f62a5504cf254a0058d35a058301", // 签名7 // 签名样本7
		"030dbbfdd29fd0f07a5ff7fab7d655c69663b8b028132628469ce33f518a3a015a4aff2eba169489554c8ea87846a75f675cc66d8d0381f3d082ba979e47e2b801", // 签名8 // 签名样本8
		"af6cb3f5d3791763bd049f205a9b4d11627b9501b72dfc9d831fe86af571da262eb9eb67e1f0104cffe8c0a1e6859fa0806d3d2750e89b9b919d262db1955eff01", // 签名9 // 签名样本9
	}

	// 将私钥从16进制转换为字节数组
	privateKeyBytes, err := hex.DecodeString(privateKeyHex) // 把私钥的16进制字符串转换为字节数组 // 解析私钥字符串为字节数组
	if err != nil {                                         // 判断是否出错 // 检查解码是否成功
		log.Fatal(err) // 出错则记录日志并终止程序 // 错误处理：终止程序并打印错误
	}

	// 将字节数组转换为私钥对象
	privateKey, err := crypto.ToECDSA(privateKeyBytes) // 将字节数组转换为 ECDSA 私钥对象 // 从字节数组生成ECDSA私钥
	if err != nil {                                    // 判断是否出错 // 检查私钥转换是否成功
		log.Fatal(err) // 出错则记录日志并终止程序 // 错误处理：终止程序并打印错误
	}

	// 对两个消息进行哈希
	hashOne := crypto.Keccak256Hash([]byte(messageOne)) // 对消息1计算 Keccak256 哈希 // 计算第一个消息的Keccak256哈希值
	hashTwo := crypto.Keccak256Hash([]byte(messageTwo)) // 对消息2计算 Keccak256 哈希 // 计算第二个消息的Keccak256哈希值

	// 使用私钥对哈希进行签名
	signatureOne, err := crypto.Sign(hashOne.Bytes(), privateKey) // 用私钥对消息1的哈希值进行签名 // 对第一个消息的哈希签名
	if err != nil {                                               // 判断是否出错 // 检查签名操作是否成功
		log.Fatal(err) // 出错则记录日志并终止程序 // 错误处理：终止程序并打印错误
	}

	signatureTwo, err := crypto.Sign(hashTwo.Bytes(), privateKey) // 用私钥对消息2的哈希值进行签名 // 对第二个消息的哈希签名
	if err != nil {                                               // 判断是否出错 // 检查签名操作是否成功
		log.Fatal(err) // 出错则记录日志并终止程序 // 错误处理：终止程序并打印错误
	}

	// 将签名转换为16进制字符串格式
	signatureOneHex := hex.EncodeToString(signatureOne) // 将第一个签名转换为16进制字符串 // 转换第一份签名为16进制字符串
	signatureTwoHex := hex.EncodeToString(signatureTwo) // 将第二个签名转换为16进制字符串 // 转换第二份签名为16进制字符串

	// 遍历签名数组，寻找与两个消息签名匹配的签名
	var foundSignatureOne, foundSignatureTwo string // 声明两个变量用于存储匹配到的签名 // 声明变量用于保存匹配成功的签名
	for _, sig := range sigArray {                  // 遍历预存的签名数组 // 循环遍历所有预设签名
		if sig == signatureOneHex { // 如果签名与第一个消息签名匹配 // 判断是否匹配第一个消息的签名
			foundSignatureOne = sig // 保存匹配的签名 // 记录匹配到的第一个签名
		}
		if sig == signatureTwoHex { // 如果签名与第二个消息签名匹配 // 判断是否匹配第二个消息的签名
			foundSignatureTwo = sig // 保存匹配的签名 // 记录匹配到的第二个签名
		}
		if foundSignatureOne != "" && foundSignatureTwo != "" { // 如果两个签名都已匹配成功，退出循环 // 检查是否两份签名都已找到
			break // 退出循环 // 结束循环
		}
	}

	// 返回找到的两个签名
	return foundSignatureOne, foundSignatureTwo // 返回两个匹配的签名字符串 // 返回结果：匹配到的两个签名
}

// 主函数
func main() { // 程序入口函数 // 主函数，程序从这里开始执行
	sig1, sig2 := searchSignature()  // 调用 searchSignature 函数获取签名 // 调用函数并接收返回的签名
	fmt.Printf("找到的签名1: %s\n", sig1) // 打印第一个匹配的签名 // 打印第一份签名
	fmt.Printf("找到的签名2: %s\n", sig2) // 打印第二个匹配的签名 // 打印第二份签名
}

/* 这个程序就像是在“比对签名”：
它先用一把私钥对两条消息盖上章（签名），然后去一堆已有的“签名记录”里挨个查，看看有没有正好盖过的一样的章。
找到了，就说明这些消息确实是这把私钥“亲手签”的，真假一目了然。 */
