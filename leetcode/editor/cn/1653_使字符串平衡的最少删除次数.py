# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def minimumDeletions(self, s: str) -> int:
        leftb, righta = 0, s.count('a')
        res = righta
        for c in s:
            if c == 'a':
                righta -= 1
            else:
                leftb += 1
            res = min(res, leftb + righta)
        return res



# leetcode submit region end(Prohibit modification and deletion)

# pytest leetcode/editor/cn/1653_使字符串平衡的最少删除次数.py
def test_minimum_deletions():
    assert Solution().minimumDeletions("aababbab") == 2
    # assert Solution().minimumDeletions("aaabbab") == 3
