package MyString

/**
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/

/**
 * 动态规划解法   假设字符串k是原字符串s的i-j之间的字符串即k=s[i,j]
 * k有两种可能 是回文 不是回文（这个又包括两种情况不是回文，当j<i时，s[i,j]不合法）
 * 边界问题字符在大于2之上。一个字符串或者两个字符串相同则为回文串且j-i+1的最大值就是我们所求的长字符串
 */
func LongestPalindrome(s string) string {
	var StrLen = len(s)
	if StrLen < 2 {
		return s
	}
	for i := 0; i < StrLen; i++ {

	}

}

/**
* 是否是回文字符串
 */
func isPalindromeNew(s string) bool {
	var len = len(s)
	if len <= 0 {
		return false
	}
	var new_string = ""
	for i := len - 1; i >= 0; i-- {

		if isalnum(s[i]) == true {
			new_string += string(s[i])
		}

	}
	return new_string == s
}

/**
* 判断是否数字或英文字母
 */
func isNumChar(ch byte) bool {

	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
