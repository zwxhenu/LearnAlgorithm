package MyString

/**
题目描述：
	给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。

示例 1:

	输入: ["abcd","dcba","lls","s","sssll"]
	输出: [[0,1],[1,0],[3,2],[2,4]]
	解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2:

	输入: ["bat","tab","cat"]
	输出: [[0,1],[1,0]]
	解释: 可拼接成的回文串为 ["battab","tabbat"]
解题思路：


*/

import (
	"regexp"
	"strings"
)

/**
* 判断字符串是否是回文
 */
func IsPalindrome(s string) bool {
	var filter_string = filterContent(s)
	var len = len(filter_string)
	if len <= 1 {
		return false
	}
	// 判断字符串子串是否为字母或者数字
	var reg = regexp.MustCompile(`^[A-Za-z0-9]+$`)
	// 遍历反转字符串比较原来字符串相同则是回文，否则不是
	var new_string = ""
	for i := len - 1; i >= 0; i-- {
		// 验证字符是否是字母或者数字
		//var is_number_char = reg.FindAllString(string(filter_string[i]), -1)
		//var is_number_char = reg.FindAllStringSubmatch(string(filter_string[i]), -1)
		var is_number_char = reg.FindString(string(filter_string[i]))
		//var find_char_index = reg.FindStringIndex(string(filter_string[i]));
		if is_number_char != "" {
			new_string += string(filter_string[i])
		}
	}
	return new_string == filter_string
}

/**
* 双指针解法
 */
func PointerIsPalindrome(s string) bool {
	var filter_string = filterContent(s)
	var len = len(filter_string)
	if len <= 1 {
		return false
	}
	var left_string = ""
	var counter = 0
	var rigth_string = ""
	for i := 0; i < len; i++ {
		counter++
		k := len - 1 - i
		left_string += string(filter_string[i])
		rigth_string += string(filter_string[k])
		if counter == i && left_string == rigth_string {
			return true
		}
	}
	return false
}

/**
* 判断字符串是否是字母或者数字
 */
func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func filterContent(s string) string {
	var reverse_str = s
	if len(s) <= 0 {
		return ""
	}
	filter_char := [8]string{" ", ",", ", ", ";", "。", "?", ".", ":"}
	for i := range filter_char {
		reverse_str = strings.Replace(reverse_str, filter_char[i], "", -1)
	}
	return strings.ToLower(reverse_str)
}
