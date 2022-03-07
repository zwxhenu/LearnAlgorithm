package MyArray

import "strings"

/**
题目描述：

	给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

	假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

	注意：每次拼写（指拼写词汇表中的一个单词）时，chars 中的每个字母都只能用一次。

	返回词汇表 words 中你掌握的所有单词的 长度之和。


示例 1：

	输入：words = ["cat","bt","hat","tree"], chars = "atach"
	输出：6
	解释： 可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。

示例 2：

	输入：words = ["hello","world","leetcode"], chars = "welldonehoneyr"
	输出：10
	解释：可以形成字符串 "hello" 和 "world"，所以答案是 5 + 5 = 10。


提示：
	1 <= words.length <= 1000
	1 <= words[i].length, chars.length <= 100
	所有字符串中都仅包含小写英文字母

解题思路：
	（1）哈希表法 分别给每个单词和词库进行构造字母的哈希数组，然后遍历单词字母哈希数组中的每个字母和词库的字母进行比较，大于的则不符合条件
	（2）不穷举字母法 遍历单词中的每个字母，统计单词中某个字母的总数是否大于词库中该字母总数，如果大于则不符合跳过进行下个单词的验证，否则对单词长度进行累加


*/

/**
hash表的解法
*/
func CountCharacters(words []string, chars string) int {
	if len(words) == 0 {
		return 0
	}
	// 构建字符哈希表
	chs := make([]byte, 'z'-'A'+1)
	for _, ch := range chars {
		chs[ch-'A']++
	}

	// 初始化长度之和
	res := 0

	// 遍历字符串数组
	for _, str := range words {
		if isContained(str, chs) { // 如果包含
			res = res + len(str)
		}
	}
	return res
}

func isContained(str string, chs []byte) bool {
	ws := make([]byte, 'z'-'A'+1)
	for _, w := range str {
		ws[w-'A']++
	}
	// --比较两个哈希表的元素
	for i := 0; i < len(ws); i++ {
		if len(ws) > len(chs) {
			return false
		}
		if ws[i] > chs[i] {
			return false
		}
	}
	return true
}

/**
不穷举的解法
*/
func NotExhaustiveAttackMethod(words []string, chars string) int {
	// chars中的字母只能用一次
	var length int = 0
outloop:
	for _, val := range words {
		// 取每一个单词
		for _, value := range val {
			// 取单词的一个字母
			if strings.Count(val, string(value)) > strings.Count(chars, string(value)) {
				// 如果在单词中某个字母的数量大于词库中的数量，说明拼写不成
				continue outloop
			}
		}
		// 循环走完，说明这个单词里的字母都在chars中
		length += len(val)
	}
	return length
}
