package main
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
import(
	"fmt"
)

func longestPalindrome(s string) string{

}
/**
* 是否是回文字符串
*/
func isPalindrome(s string) bool{
	var len = len(s)
	if len <= 0{
		return false
	}
	var new_string = ""
	for i:=len-1;i>=0;i--{

		if isalnum(s[i]) == true{
			new_string +=string(s[i])
		}
		
	}
	return new_string == s
}
/**
* 判断是否数字或英文字母
*/
func isNumChar(ch byte) bool{

	return (ch >='a' && ch <='z') || (ch >='A' && ch <= 'Z') || (ch >='0' && ch <= '9')
}

func main(){

	fmt.Println("hello")
}