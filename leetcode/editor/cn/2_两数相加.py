# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


from typing import Optional


class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if l1 is None:
            return l2
        if l2 is None:
            return l1
        carry = 0
        head = ListNode(0)
        cur = head
        while l1 is not None or l2 is not None or carry != 0:
            x = (l1.val if l1 is not None else 0) + (l2.val if l2 is not None else 0) + carry
            carry = x // 10
            cur.next = ListNode(x % 10)
            cur = cur.next
            if l1 is not None:
                l1 = l1.next
            if l2 is not None:
                l2 = l2.next
        return head.next
# leetcode submit region end(Prohibit modification and deletion)
