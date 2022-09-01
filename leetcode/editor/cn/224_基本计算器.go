package leet

//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
//
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
//
//
// 示例 1：
//
//
//输入：s = "1 + 1"
//输出：2
//
//
// 示例 2：
//
//
//输入：s = " 2-1 + 2 "
//输出：3
//
//
// 示例 3：
//
//
//输入：s = "(1+(4+5+2)-3)+(6+8)"
//输出：23
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 3 * 10⁵
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
// '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
// '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
// 输入中不存在两个连续的操作符
// 每个数字和运行的计算将适合于一个有符号的 32位 整数
//
// Related Topics栈 | 递归 | 数学 | 字符串
//
// 👍 818, 👎 0
//
//
//
//

//leetcode submit region begin(Prohibit modification and deletion)
func calculate(s string) int {
	stack := []int{}
	// sign  1 --> +
	// sign -1 --> -
	// 默认第一个字符默认携带 '+' , sign 指定为 1
	res, num, sign := 0, 0, 1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // 计算出字符串出数字值，包含多个字符: 7 17
			num = num*10 + int(s[i]-'0')
		case '+': // 将之前的总和，公式的左半部分 res 和 公式的右半部分num通过 符号 sign 累计起来
			res += sign * num
			num, sign = 0, 1
		case '-':
			res += sign * num
			num, sign = 0, -1
		case '(': // 将之前的总和作为公式左边部分res 以及符号 sign 推入栈, 从新初始化后进行计算
			stack = append(stack, sign)
			stack = append(stack, res)
			res, num, sign = 0, 0, 1
		case ')': // 将表达式总和 作为公式的右边部分 赋值给 num，从stack 取出 res 和 sign
			num = res + sign*num
			res = stack[len(stack)-1]
			sign = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}
	}
	if num != 0 { // 还会有负数的情况
		res += sign * num
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
