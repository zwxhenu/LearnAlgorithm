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
	if len(s) < 2 {
		return s
	}

	// dp[i][j]表示字符串s[i,j]是否为回文串
	dp := make([][]bool, len(s))

	//初始化列
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}
	/*
		  4  e           *
		  3  d        *  *
		  2  c     *  *  *
		  1  b  *  *  *  *
		  0  a  b  c  d  e
			 0	1  2  3  4
	*/
	start := 0  //最大回文串的开始下标
	maxLen := 0 //最大回文串的长度-1(最后再+1就可以表示长度,不要在循环里面每次都+1)
	// i要从大到小
	for i := len(s) - 1; i >= 0; i-- {
		// j要从小到大
		for j := i; j < len(s); j++ {
			length := j - i
			// 可以写成一行 dp[i][j] = s[i] == s[j] && (length <= 2 || dp[i+1][j-1])
			if s[i] == s[j] {
				if length <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && length > maxLen {
				maxLen = length
				start = i
			}
		}
	}
	// for i, v := range dp {
	// 	fmt.Printf("dp[%d]=%+v\n", i, v)
	// }
	return s[start : start+maxLen+1]
}

/**
* 暴力破解法 效率太低
 */
func NormalGetLongestPalindString(s string) string {
	var StrLen = len(s)
	if StrLen <= 1 {
		return s
	}
	var verify = ""
	var MaxString = ""
	for i := 0; i < StrLen; i++ {
		for j := i + 1; j <= StrLen; j++ {
			verify = s[i:j]
			if IsPalindromeNew(verify) {
				if len(verify) > len(MaxString) {
					MaxString = verify
				}
			}
		}
	}
	return MaxString
}

/**
* 是否是回文字符串
 */
func IsPalindromeNew(s string) bool {
	var len = len(s)
	if len <= 0 {
		return false
	}
	var new_string = ""
	for i := len - 1; i >= 0; i-- {

		if isNumChar(s[i]) == true {
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
