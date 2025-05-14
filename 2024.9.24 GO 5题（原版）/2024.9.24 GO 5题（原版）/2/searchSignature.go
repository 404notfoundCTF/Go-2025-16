package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func searchSignature() (string, string) {
	privateKeyHex := "6e0aac9f7bf63010390545b055c23505259a2ce4c87ddbf3acf48cefd83a6dac"
	messageOne := "hello world"
	messageTwo := "hello blockchain"
	sigArray := []string{
		"3cc4af10427f9b0cdb196fa3542ccc1f90ae11f5296ac666bea0f44f935480ff7ab212d01a1c1e5245a7a099f4b64e2de47008f8626d4fcb172facc5e087dd3b00",
		"d633b72ee1f04e70a7d9043e6c8bcc8b64920ffb24a8dcd27ce487741c94442b2cc3899d553734550e51eb6b6755861578397b892d5308497d4f7c787b9752a700",
		"458be524741b15bd3c665f1ef5d9564ec95935adf0957606db11759bf5e162ab5b53f63d490e91292c61250f6cf4ee96460090a79ecc21f95e5bb7b04a2e793900",
		"aa17c15fd8618ec2623c776214d6e2606d7f80574b5af026fe8c299fad52364668b761293e70cdc28530e878cc8e94e12981f24de98853c58f32929c4864c5c801",
		"bf8138c93e565985c4d48dbe00ddc45c12c30ecc862290cb009ea452b3faf589325a2454c739cda4180d3c7bc1f63ec7e529636a4b1815d2c252c93e14281cdf01",
		"a73ad2c912b1d59827bccd70042f74b403eaad6c1dcad2e4d2a01166eb7e6d5c5a0669739e7380270d888627e5074453413a6b224d038e5d10717f56555c71d900",
		"75c362d664e8f79a1cdf4524b5ae8dec1a4541500b5ab5738e093000e160e79025cb103b56f5d6d28bac627d0cee0507b3c0f1f62a5504cf254a0058d35a058301",
		"030dbbfdd29fd0f07a5ff7fab7d655c69663b8b028132628469ce33f518a3a015a4aff2eba169489554c8ea87846a75f675cc66d8d0381f3d082ba979e47e2b801",
		"af6cb3f5d3791763bd049f205a9b4d11627b9501b72dfc9d831fe86af571da262eb9eb67e1f0104cffe8c0a1e6859fa0806d3d2750e89b9b919d262db1955eff01",
	}
	//以下注释的内容都要掌握补全↓

	// 将私钥从16进制转换为字节数组
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// 将字节数组转换为私钥对象
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// 对两个消息进行哈希
	hashOne := crypto.Keccak256Hash([]byte(messageOne))
	hashTwo := crypto.Keccak256Hash([]byte(messageTwo))

	// 使用私钥对哈希进行签名
	signatureOne, err := crypto.Sign(hashOne.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	signatureTwo, err := crypto.Sign(hashTwo.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 将签名转换为16进制字符串格式
	signatureOneHex := hex.EncodeToString(signatureOne)
	signatureTwoHex := hex.EncodeToString(signatureTwo)

	// 遍历签名数组，寻找与两个消息签名匹配的签名
	var foundSignatureOne, foundSignatureTwo string
	for _, sig := range sigArray {
		if sig == signatureOneHex {
			foundSignatureOne = sig
		}
		if sig == signatureTwoHex {
			foundSignatureTwo = sig
		}
		if foundSignatureOne != "" && foundSignatureTwo != "" {
			break
		}
	}

	// 返回找到的两个签名
	return foundSignatureOne, foundSignatureTwo
}

// 主函数
func main() {
	sig1, sig2 := searchSignature()
	fmt.Printf("找到的签名1: %s\n", sig1)
	fmt.Printf("找到的签名2: %s\n", sig2)
}
