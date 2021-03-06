package mleetcode

import (
	"fmt"
	"strings"
	"testing"
)

//leetcode 3
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//滑动窗口

func lengthOfLongestSubstring(s string) int {
	var length int
	var str string
	left := 0
	right := 0
	//当前最长的字符串
	str = s[left:right]
	for right = 0; right < len(s); right++ {
		if index := strings.IndexByte(str, s[right]); index != -1 {
			//存在
			left = left + index + 1
		}
		str = s[left : right+1]
		if len(str) > length {
			length = len(str)
		}
	}
	return length
}

func TestString(t *testing.T) {
	aaa := lengthOfLongestSubstring("asadbascs")
	fmt.Println(aaa)
}
