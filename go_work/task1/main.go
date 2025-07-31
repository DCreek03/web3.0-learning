package main

import (
	"fmt"
	"slices"
)

func main() {
	// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
	// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
	// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	fmt.Println("数组中只出现一次元素的是：", findSameValue([]int{1, 2, 1, 2, 6}))

	// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	fmt.Println("给定的整数是否是回文的结果是：", IsCircle(1221))

	// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	fmt.Println("判断给定的字符串的结果是:", isValid("({})"))

	// 编写一个函数来查找字符串数组中的最长公共前缀。
	// 如果不存在公共前缀，返回空字符串 ""。
	fmt.Println("存在的最长的公共前缀是:", maxSameHead([]string{"flower", "flow", "flight"}))

	// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
	// 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	// 将大整数加 1，并返回结果的数字数组。
	fmt.Println("加一的结果是：", plusOne([]int{1, 2, 3, 9}))

	// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，
	// 使每个元素 只出现一次 ，返回删除后数组的新长度。
	// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	fmt.Println("删除重复出现的元素结果是：", deleteSame([]int{1, 1, 2, 2, 3, 3, 9}))

	// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
	fmt.Println("删除重复出现的元素结果是：", merge([][]int{{1, 3}, {2, 6}, {20, 32}}))

	// 给定一个整数数组 nums 和一个整数目标值 target，
	// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
	fmt.Println("两数之和的结果是:", twoSum2([]int{1, 1, 3}, 4))

}
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {

		if i, ok := m[target-value]; ok {
			return []int{i, index}
		}
		m[value] = index
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return make([]int, 0)
// }

func merge(intervals [][]int) (result [][]int) {
	// 左端点排序
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	for _, p := range intervals {
		m := len(result)
		if m > 0 && p[0] <= result[m-1][1] { // 左端点跟上一个右端点进行比较
			result[m-1][1] = max(result[m-1][1], p[1])
		} else {
			result = append(result, p)
		}
	}
	return

}

func deleteSame(x []int) int {
	// 初始化 k=1，表示保留的元素要填入的下标。
	// 从 i=1 开始遍历 nums。
	// 如果 nums[i]=nums[i−1]，那么 nums[i] 是重复项，不保留。
	// 如果 nums[i]=nums[i−1]，那么 nums[i] 不是重复项，保留，填入 nums[k] 中，然后把 k 加一。
	// 遍历结束后，k 就是 nums 中的唯一元素的数量，返回 k。
	k := 1
	for i := 1; i < len(x); i++ {
		if x[i] != x[i-1] {
			x[k] = x[i]
			k++
		}
	}
	return k
}

func plusOne(x []int) []int {
	if len(x) == 0 {
		return []int{1}
	}
	sum := 0
	mult := 10
	for i := len(x) - 1; i >= 0; i-- {

		if i == len(x)-1 {
			sum += x[i]
		} else {
			sum += x[i] * mult
			mult *= 10
		}
	}
	sum += 1
	result := []int{}
	for sum > 10 {
		result = append(result, sum%10)
		sum /= 10
	}
	result = append(result, sum)
	slices.Reverse(result)
	return result

}

func maxSameHead(strs []string) string {
	// 具体算法如下：
	// 从左到右遍历 strs 的每一列。设当前遍历到第 j 列，从上到下遍历这一列的字母。
	// 设当前遍历到第 i 行，即 strs[i][j]。如果 j 等于 strs[i] 的长度，
	// 或者 strs[i][j]=strs[0][j]，说明这一列的字母缺失或者不全一样，那么最长公共前缀的长度等于 j，
	// 返回 strs[0] 的长为 j 的前缀。
	// 如果没有中途返回，说明所有字符串都有一个等于 strs[0] 的前缀，那么最长公共前缀就是 strs[0]。

	s0 := strs[0]
	for j, c := range s0 {
		for _, s := range strs {
			if j == len(s) || s[j] != byte(c) {
				return s0[:j]
			}
		}
	}
	return s0

}

func isValid(s string) bool {
	if len(s)%2 != 0 { // s 长度必须是偶数
		return false
	}

	mp := map[rune]rune{')': '(', ']': '[', '}': '{'}
	st := []rune{}
	for _, c := range s {
		if mp[c] == 0 { // c 是左括号
			st = append(st, c) // 入栈
		} else { // c 是右括号
			if len(st) == 0 || st[len(st)-1] != mp[c] {
				return false // 没有左括号，或者左括号类型不对
			}
			st = st[:len(st)-1] // 出栈
		}
	}
	return len(st) == 0 // 所有左括号必须匹配完毕
}

func IsCircle(i int) bool {
	if i < 0 {
		return false
	}
	// 使用切片 辗转相除
	digits := []int{}
	for i > 10 {
		digits = append(digits, i%10)
		i /= 10
	}
	// 把个位数添加进去
	digits = append(digits, i)
	for i, j := 0, len(digits)-1; i < len(digits)-1/2; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}

func findSameValue(v []int) int {
	// 定义一个map 需要注意的是这里需要对map进行初始化
	m := make(map[int]int)
	for i := range v {
		count, flag := m[v[i]]
		if flag {
			m[v[i]] = count + 1
		} else {
			m[v[i]] = 1
		}
	}

	for index, count := range m {
		if count == 1 {
			return index
		}
	}
	return 0
}
