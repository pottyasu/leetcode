/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	var result []int
	hashmap := make(map[int]int)

	for i, v1 := range nums {
		comp := target - v1
		if _, ok := hashmap[comp]; ok {
			result = append(result, hashmap[comp], i)
			break
		}
		hashmap[v1] = i
	}

	return result
}

// @lc code=end

