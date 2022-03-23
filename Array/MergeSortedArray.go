package MyArray

import "fmt"

/**
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。



示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
示例 3：

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。


提示：

nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[j] <= 109


进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？

*/

// 逆向双指针
func ReverseTwoPointArrayMerge(nums1 []int, n int, nums2 []int, m int) {

	for p1, p2, tail := n-1, m-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
	fmt.Println(nums1)
}

//双指针 有序数组，从头部取出数据比较最小的放到新的数组中
func TwoPointerArrayMerge(nums1 []int, n int, nums2 []int, m int) {
	var newArr = make([]int, 0, m+n)
	var p1, p2 = 0, 0
	for {
		// 判断指针结束的时候
		if p1 == n {
			newArr = append(newArr, nums2[p2:]...)
			break
		}
		if p2 == m {
			newArr = append(newArr, nums1[p1:]...)
			break
		}
		if nums1[p1] > nums2[p2] {
			newArr = append(newArr, nums2[p2])
			p2++
		} else if nums1[p1] == nums2[p2] {
			newArr = append(newArr, nums1[p1])
			newArr = append(newArr, nums2[p2])
			p1++
			p2++
		} else {
			newArr = append(newArr, nums1[p1])
			p1++
		}
	}
	fmt.Println(newArr)
}

// 冒泡排序解决
func NewDealArrayMerge(nums1 []int, n int, nums2 []int, m int) (result []int) {
	copy(nums1[n:], nums2)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			if nums1[i] == 0 {
				continue
			}
			if nums1[i] < nums1[j] {
				tmp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = tmp
			}
		}
	}
	return nums1
}

func DealArrayMerge(nums1 []int, n int, nums2 []int, m int) (result []int) {
	if n <= 0 && m <= 0 {
		return make([]int, m+n)
	}
	if n <= 0 {
		return nums2
	}
	if m <= 0 {
		return nums1
	}
	var newArr = make([]int, n+m)
	var firstIndex = 0
	for i := 0; i < m; i++ {
		newArr[i] = nums2[i]
		firstIndex = i
	}
	for i := 0; i < n; i++ {
		firstIndex++
		if firstIndex < 6 {
			newArr[firstIndex] = nums1[i]
		}
	}
	for i := 0; i < len(newArr); i++ {
		for j := 0; j < len(newArr); j++ {
			if newArr[i] == 0 {
				continue
			}
			if newArr[i] < newArr[j] {
				tmp := newArr[i]
				newArr[i] = newArr[j]
				newArr[j] = tmp
			}
		}
	}
	return newArr
}
