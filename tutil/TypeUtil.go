package tutil

import "strconv"

//  author: laoniqiu
//  since: 2022/8/22
//  desc: tutil

type typeUtil struct{}

var typeUtilInstance *typeUtil

func GetInstanceByTypeUtil() *typeUtil {
	return typeUtilInstance
}

func (*typeUtil) StrToInt(str string) int {
	s, _ := strconv.ParseFloat(str, 10)
	return int(s)
}
