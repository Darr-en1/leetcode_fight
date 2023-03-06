# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = 0
        _len = 0
        hashset = set()
        for idx, val in enumerate(s):
            while val in hashset:
                hashset.remove(s[start])
                start += 1
            hashset.add(val)
            _len = max(_len, idx - start + 1)
        return _len


# leetcode submit region end(Prohibit modification and deletion)

if __name__ == '__main__':
    assert Solution().lengthOfLongestSubstring("abcabcbb") == 3
