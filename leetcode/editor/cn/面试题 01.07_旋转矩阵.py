# leetcode submit region begin(Prohibit modification and deletion)
from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        for i in range(len(matrix)):
            for j in range(i, len(matrix)):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

            # leetcode submit region end(Prohibit modification and deletion)
        for row in matrix:
            reverse(row)


def reverse(arr: List[int]):
    i, j = 0, len(arr) - 1
    while j > i:
        arr[i], arr[j] = arr[j], arr[i]
        i += 1
        j -= 1


if __name__ == '__main__':
    array = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12], [13, 14, 15, 16]]
    Solution().rotate(array)
    print(array)

    array1 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    Solution().rotate(array1)
    print(array1)
