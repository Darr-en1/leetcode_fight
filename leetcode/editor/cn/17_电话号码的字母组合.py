# leetcode submit region begin(Prohibit modification and deletion)
import itertools
from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return list()

        mapping = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz"
        }

        def inner(index):
            if index == len(digits):
                combinations.append("".join(combination))
            else:
                digit = digits[index]
                for letter in mapping[digit]:
                    combination.append(letter)
                    inner(index + 1)
                    combination.pop()

        combination = []
        combinations = []
        inner(0)
        return combinations

    def letterCombinations1(self, digits: str) -> List[str]:
        mapping = {
            "2": "abc",
            "3": "def",
            "4": "ghi",
            "5": "jkl",
            "6": "mno",
            "7": "pqrs",
            "8": "tuv",
            "9": "wxyz"
        }
        # 笛卡尔积
        return ["".join(combination) for combination in
                itertools.product(*[mapping[digit] for digit in digits])] if digits else []


# leetcode submit region end(Prohibit modification and deletion)


def test_letterCombinations():
    assert Solution().letterCombinations("23") == ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
    assert Solution().letterCombinations("") == []
