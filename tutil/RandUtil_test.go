package tutil

import (
	"strings"
	"testing"
)

func TestRandNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(RandNumber(100))
	}
	s := strings.ToUpper(RandCharacterString(320))
	t.Log(s)

	var sum uint64 = 36
	for i := 0; i < 16; i++ {
		sum *= 36
	}
	t.Log(sum)

	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = int(RandNumber(100))
	}
	t.Log(arr)
	//quickSort(arr)
	quick_sort(arr, 0, len(arr)-1)
	t.Log(arr)
}

func quickSort(s []int) {
	length := len(s)
	if length < 2 {
		return
	}
	head, tail := 0, length-1
	value := s[head] // 标尺元素
	for head < tail {
		if s[head+1] > value {
			// 大于标尺元素，把元素换到最右边
			s[head+1], s[tail] = s[tail], s[head+1]
			tail--
		} else if s[head+1] < value {
			//小于标尺元素，换位置并把标尺右移动一位。
			s[head], s[head+1] = s[head+1], s[head]
			head++
		} else {
			//相等不交换
			head++
		}
	}
	//标尺左边的元素都小于等于标尺元素（s[head]），右边的元素大于等于标尺元素。
	quickSort(s[:head])
	quickSort(s[head+1:])
}

func partition(arr []int, l int, r int) int {
	var i, j int = l - 1, r + 1
	var mid int = (l + r) / 2

	x := arr[mid]
	for i < j {
		i++
		for arr[i] < x {
			i++
		}
		j--
		for arr[j] > x {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return j
}

func quick_sort(arr []int, l int, r int) {
	if r <= 0 || l >= r {
		return
	}
	stack := []int{}
	i := 0
	j := 0
	stack = append(stack, r)
	stack = append(stack, l)

	for len(stack) > 0 {
		i = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		j = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i < j {
			k := partition(arr, i, j)
			if k > i {
				stack = append(stack, k)
				stack = append(stack, i)
			}
			if k < j {
				stack = append(stack, j)
				stack = append(stack, k+1)
			}
		}
	}
}
