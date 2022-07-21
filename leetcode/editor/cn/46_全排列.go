package leet

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
// Related Topics数组 | 回溯
//
// 👍 2117, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)

func permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}
	track := []int{}
	numsLen := len(nums)

	var dfs func()
	dfs = func() {
		if len(track) == numsLen {
			temp := make([]int, numsLen)
			// 被拷贝的字符串在后面
			copy(temp, track)
			res = append(res, temp)
		}
		for _, num := range nums {
			if visited[num] {
				continue
			} else {
				visited[num] = true
				track = append(track, num)
				dfs()
				delete(visited, num)
				track = track[:len(track)-1]
			}
		}
	}
	dfs()
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
