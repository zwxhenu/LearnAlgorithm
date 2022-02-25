package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true
 

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成

*/

import(
	"fmt"
)
/**
 *	栈思路解法
 */
func isValid(s string)bool{
	var str_len = len(s)
	if str_len%2 != 0{
		return false
	}
	var match_str = map[byte]byte{
			')':'(',
			']':'[',
			'}':'{',
		}
	var array = make([]byte, 0)
	// 字符串转字节数组
	var data []byte = []byte(s)
	for i:=0;i< str_len; i++ {
		// 存在闭合符的进行判断
		if match_str[data[i]] > 0 {
			// 栈为空或者栈顶字符不等于匹配字符 存在不闭合字符串
			if len(array) == 0 || array[len(array)-1] != match_str[data[i]] {		
				return false
			}
			// 如果相同 则出栈
			array = array[:len(array)-1]
		}else{// 不存在闭合符的入栈
			array = append(array, data[i])
		}
	}
	
	return len(array) == 0
}

func main() {
	
	var str = "(){{}{}[]"
	var result = isValid(str)
	fmt.Println(result)
}
