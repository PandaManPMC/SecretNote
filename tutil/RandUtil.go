package tutil

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var chaArr = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

//	生成真随机数 0-9 crypto/rand
//	num int 随机数长度
func RandCrypto(num int) string {
	str := strings.Builder{}
	for i := 0; i < num; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(10))
		str.WriteString(fmt.Sprintf("%d", result))
	}
	return str.String()
}

//	RandCharacterString 生成 0-9 a-z 随机数
//	num int 随机数长度
func RandCharacterString(num int) string {
	str := strings.Builder{}
	max := len(chaArr)
	for i := 0; i < num; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
		str.WriteString(chaArr[result.Int64()])
	}
	return str.String()
}

//	生成真随机数数字 max 以内
//	max int 随机数最大
func RandNumber(max int) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return result.Int64()
}
