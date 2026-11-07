package homework01

import "sort"

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 相同的数字异或为 0，成对抵消后剩下的就是只出现一次的元素
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// 负数不是回文；除 0 以外个位为 0 的数也不是（如 10 反转后是 01）
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 只反转后一半数字，避免整数溢出
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 位数为偶数时 x == reversed；为奇数时中间那一位留在 x 上
	return x == reversed || x == reversed/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0, len(s))

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			// 右括号必须与栈顶左括号配对，栈为空说明缺少左括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false // 只允许出现括号
		}
	}

	// 全部处理完后栈为空，说明每个左括号都匹配上了
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准，逐位比对其它字符串
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// 长度不够或字符不同，前缀到此为止
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 当前位满十进位，本位归零
	}

	// 每一位都是 9（如 999），需要在最前面补一个 1
	return append([]int{1}, digits...)
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 慢指针指向下一个写入位置，快指针负责往后扫描
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 按起点升序排序，这样只需和上一个合并结果比较
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{{intervals[0][0], intervals[0][1]}}
	for _, cur := range intervals[1:] {
		last := merged[len(merged)-1]
		if cur[0] <= last[1] { // 与上一个区间重叠，合并到最后一个
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else { // 不重叠，开启一个新区间
			merged = append(merged, []int{cur[0], cur[1]})
		}
	}

	return merged
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 记录已遍历数字的下标，用补数去查找
	seen := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil // 题目保证有解
}
