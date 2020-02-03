package main

import (
	"fmt"
	"math"
)

/**
LeetCode69 如果求一个数的平方根？
 */

func main() {

	for i := 1; i <= 1000; i++ {
		fmt.Println(i, mySqrt(i))
	}
}

func mySqrt(target int) int {
	if target < 1 {
		return 0
	}
	if target == 1 {
		return 1
	}
	var low, high int = 0, target
	var mid int = (high + low) / 2
	for high - low > 1 {
		if mid * mid > target {
			high = mid
		} else if mid * mid < target {
			low = mid
		} else {
			break
		}
		mid = (high + low) / 2
	}
	return mid
}

/**
二分法,平方保留6位小数精度
存在的问题：
	1.对于某些精确平方，结果反而存在一定误差，如：9，16
 */
func BinarySearchSqrt(target float64) float64 {
	if target < 0 {
		return -1
	}
	var low, high float64 = 0, target
	var mid float64 = (high + low) / 2
	for high - low >= 1.e-6 {
		if mid * mid > target {
			high = mid
		} else if mid * mid < target {
			low = mid
		} else {
			break
		}
		mid = (high + low) / 2
	}
	return mid
}

/**
牛顿迭代法计算平方根
求解思路，参考【./牛顿迭代法求平方根.jpg】
 */
func NewtonCalculateSqrt(target float64) float64 {
	var x1, x2 float64 = 0, target
	for math.Abs(x1 - x2) > 1.e-6 {
		x1 = x2
		x2 = x1 / 2 + target / (2 * x1)
	}
	return x1
}