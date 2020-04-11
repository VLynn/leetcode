/*
 * 返回最长的回文子串
 * 遍历字符串，以每个字符和字符右边空隙为对称中心，判断左右两边是否对称
 * 时间复杂度O(n2)，空间复杂度O(1)
 */

func longestPalindrome(s string) string {
    // 最长子串左右索引
    longLeft, longRight := 0, 0
    // 子串左右索引
    left, right := 0, 0

    for i := 0; i < len(s); i++ {
        // 以字符和字符右边空隙为对称中心
        for k := 0; k < 2; k++ {
            left, right = i, i+k
            // 从对称中心开始，向两边遍历
            for left >= 0 && right < len(s) {
                if s[left] == s[right] {
                    left--
                    right++
                } else {
                    break
                }
            }
            // 更新最长子串
            if right - left - 2 > longRight - longLeft {
                longLeft = left + 1
                longRight = right - 1
            }
        }
    }

    // 返回子串
    if len(s) > 0 {
        return s[longLeft:longRight+1]
    } else {
        return ""
    }
}