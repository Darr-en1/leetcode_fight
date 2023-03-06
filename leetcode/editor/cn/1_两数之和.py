# leetcode submit region begin(Prohibit modification and deletion)
from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashtable = {}
        for i, v in enumerate(nums):
            if target - v in hashtable:
                return [hashtable[target - v], i]
            hashtable[v] = i
# leetcode submit region end(Prohibit modification and deletion)
