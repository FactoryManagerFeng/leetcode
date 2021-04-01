package leetcode

/**
 * 62: 一个机器人位于一个 m✖️n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 * 问总共有多少条不同的路径？
 */

/**
 *  f(m)(n) = 为走到m,n坐标总共的路径
 *  m为横坐标，n为纵坐标
 *  f(m)(n) = f(m-1)(n)+f(m)(n-1)
 *  f(m-1)(n) = f(m-2)(n)+f(m-1)(n-1)
 *  ...
 *	...
 *	f(x)(0) = 1 f(0)(x) = 1
 *	f(0)(0) = 1
 */

func HavePaths(m, n int) int {
	var f = make(map[int]map[int]int)
	for i := 0; i < m; i++ {
		f[i] = make(map[int]int)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}
