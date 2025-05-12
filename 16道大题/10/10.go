// 这题ECC公钥生成的有问题，main函数验证的时候暂时用裸ECC公钥代替
package main

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

// 判断给公钥所属算法类型，返回ECC或RSA
func IdentifyKeyType(publicKey string) string {
	// 尝试解析PEM格式
	block, _ := pem.Decode([]byte(publicKey))
	if block != nil {
		// 解析公钥内容
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err == nil {
			// 通过类型断言判断实际类型
			switch pub.(type) {
			// ========请在下方补充代码========
			// *rsa.PublicKey为RSA 公钥指针、*ecdsa.PublicKey为椭圆曲线 ECDSA 公钥指针
			case *rsa.PublicKey:
				return "RSA"
			case *ecdsa.PublicKey:
				return "ECC"
				// ========请在上方补充代码========
			}
		}
	}
	// 尝试解析裸密钥: 检查是否为32字节原始数据
	if len(publicKey) == 64 && !strings.Contains(publicKey, "-----") {
		return "ECC" // 根据假设判断
	}
	// 默认返回false根据需求调整)
	return "FALSE"
}

func main() {
	// 测试RSA公钥
	rsaPublicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvycTaLxwJ8YOG/SnT1QO
k4ED/qx83J9Gq1Xp2VGbj8Lti9jHjQ8IPoKT0j+YHc7MQTgV9MYQ6pnB/P2a/vVo
2lxJxZ4oLKgw9EoLzU6V7cLPPEjHQpOGaDtP+0vEhQgk5LGKrOUKyDp7PxgBgZnt
BZm+PjBitB1Vx0xHGQVtd9P7O/NjQF1zHQrmUxiQsV98ju3Mtr0QxOUKYuZ7cIBw
6sEwP7TYgTS6q6YoKGmRmMBRpjG8+Z1/q9RuXQFct5+P7bKuMXPtzMvqnHZ5/VlK
FmKWVARflhBj6K6d7BYJRhQzVjYuHmZZRwkRQUFvLc5ZLmM2RnLhBj3UsZCSvJPP
TQIDAQAB
-----END PUBLIC KEY-----`

	// 测试裸ECC公钥(模拟)
	rawEccKey := "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

	// 打印结果
	fmt.Println("RSA公钥类型:", IdentifyKeyType(rsaPublicKey))

	fmt.Println("裸ECC公钥类型:", IdentifyKeyType(rawEccKey))
}
