# leetcode submit region begin(Prohibit modification and deletion)
from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        start, end = 0, len(height) - 1
        res = 0
        while start < end:

            if height[start] < height[end]:
                res = max(res, height[start] * (end - start))
                start += 1
            else:
                res = max(res, height[end] * (end - start))
                end -= 1
        return res


# leetcode submit region end(Prohibit modification and deletion)

def test_maxArea():
    assert Solution().maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
    assert Solution().maxArea([1, 1]) == 1
