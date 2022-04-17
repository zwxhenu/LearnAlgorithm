package MySort

/**
* 快速排序
 */
func QuickSort(nums []int) []int {
	var numsLen = len(nums)
	if numsLen < 2 {
		return nums
	}
	var smaller = make([]int, 0)
	var bigArr = make([]int, 0)
	var targetData = nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i] < targetData {
			bigArr = append(bigArr, nums[i])
		} else {
			smaller = append(smaller, nums[i])
		}
	}
	smaller, bigArr = QuickSort(smaller), QuickSort(bigArr)
	var myArr = append(append(smaller, targetData), bigArr...)

	return myArr
}

/**
* 快速排序
 */
func NewQuickSort(nums []int) []int {

	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}
	target, index := nums[0], 1
	head, tail := 0, numsLen-1
	for head < tail {
		if nums[index] < target {
			nums[index], nums[tail] = nums[tail], nums[index]
			tail--
		} else {
			nums[index], nums[head] = nums[head], nums[index]
			index++
			head++
		}
	}
	nums[head] = target
	NewQuickSort(nums[:head])
	NewQuickSort(nums[head+1:])
	return nums
}
