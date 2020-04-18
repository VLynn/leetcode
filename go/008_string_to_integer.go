/*
 * 字符串转数字
 */

import "math"

func myAtoi(str string) int {
    res, idx, sign := 0, 0, 1

    // 跳过空格
    for idx < len(str) && str[idx] == ' ' {
        idx++
    }

    // 支持可选的正负号
    if idx == len(str) {
        return 0
    }
    if str[idx] == '-' {
        sign = -1
        idx++
    } else if str[idx] == '+' {
        idx++
    }

    // 字符转数字
    for idx < len(str) && str[idx] >= '0' && str[idx] <= '9' {
        res = res * 10 + int(str[idx] - '0')
        if sign * res > math.MaxInt32 {
            return math.MaxInt32
        }
        if sign * res < math.MinInt32 {
            return math.MinInt32
        }
        idx++
    }

    return sign * res
}