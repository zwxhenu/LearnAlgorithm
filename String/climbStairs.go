package MyString

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？



示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶


提示：

1 <= n <= 45
*/

/**
* 递归实现
 */
func ClimbStairs(n int) int {

	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

/**
* 动态规划 f(1) =1 f(2)=2 f(3)=3 f(4)=5 即f(n)=f(n-1)+f(n-2)
*
 */
func ClimbStairsDynamic(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
func ClimbStairsDynamicNew(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var ret = 0
	var pre = 2
	var ppre = 1
	for i := 3; i <= n; i++ {
		ret = pre + ppre
		ppre = pre
		pre = ret
	}
	return ret
}
