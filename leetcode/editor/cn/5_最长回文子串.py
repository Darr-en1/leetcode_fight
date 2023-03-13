# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def longestPalindrome(self, s: str) -> str:
        len_s = len(s)
        res = ""

        # 采取中间扩散的方式
        def inner(left, right):
            while left >= 0 and right < len_s and s[left] == s[right]:
                left -= 1
                right += 1
            return s[left + 1:right]  # 一定要注意，左开右闭,不合适就取上一个结果

        for i, _ in enumerate(s):
            if len(res_new := inner(i, i)) > len(res):
                res = res_new
            if len(res_new := inner(i, i + 1)) > len(res):
                res = res_new

        return res

# leetcode submit region end(Prohibit modification and deletion)
