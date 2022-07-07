package leet

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics哈希表 | 字符串 | 滑动窗口
//
// 👍 7778, 👎 0
//
//
//
//

//lengthOfLongestSubstring submit region begin(Prohibit modification and deletion)
// 正常 O(2n)  需要将map 中的数据delete
// 以下方案  O(n)  通过记录相同的位置 即可
func lengthOfLongestSubstring(s string) int {
	maxLen, start, end := 0, 0, 0
	window := make(map[byte]int)
	// golang 字符串遍历  fori  获取 value 类型为 byte(也就是uint8 类型)  forr  获取到 value 类型为 int32
	for i := 0; i < len(s); i++ {
		if idx, ok := window[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		end++

		if newLen := end - start; newLen > maxLen {
			maxLen = newLen
		}
		window[s[i]] = i
	}
	return maxLen
}
