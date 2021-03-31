package leetcode

import "math"

/**
 * 322: 你有三种硬币，2元、5元、7元，每种硬币足够多，买一本书需要27元，用最少的硬币组合
 */

/**
 *  f(x) = 最少组合类型 x为需要支付钱数
 *  有3中硬币 2,5,7
 *  f(x) = f(x-2)+1 或者 f(x) = f(x-5)+1 或者 f(x) = f(x-7)+1
 *  f(x) = f(x-y)+1
 *  ...
 *	...
 *	f(1) = 极大值
 *	f(0) = 0
 */

// coinType 硬币类型
// pay 需要支付钱数
// num 最少组合类型
func CoinChange(coinType []int, pay int) (num int64) {
	f := map[int]int64{}
	f[0] = 0
	for i := 1; i <= pay; i++ {
		// 先给每个f(x)设置一个极大值
		f[i] = math.MaxInt64
		for j := 0; j < len(coinType); j++ {
			// 当x大于钱币值且f(x-y)小于极大值且f(x)大于f(x-y)+1的时候 替换f(x)
			if i >= coinType[j] && f[i-coinType[j]] < math.MaxInt64 && f[i] > f[i-coinType[j]]+1 {
				f[i] = f[i-coinType[j]] + 1
			}
		}
	}
	return f[pay]
}
