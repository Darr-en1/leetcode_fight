# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def longestPalindrome(self, s: str) -> str:
        res = ""

        def inner(left, right):
            while left >= 0 and right < len(s) and s[left] == s[right]:
                left -= 1
                right += 1
            return s[left + 1:right]

        for i, _ in enumerate(s):
            if len(e := inner(i, i)) > len(res):
                res = e
            if len(e := inner(i, i + 1)) > len(res):
                res = e
        return res


def test_longestPalindrome():
    assert Solution().longestPalindrome("abbabb") == "bbabb"
    assert Solution().longestPalindrome("abbab") == "abba"

# leetcode submit region end(Prohibit modification and deletion)
