/*
 * 暴力字符串匹配
 * Runtime: 0 ms, faster than 100.00% of Go online submissions
 * Really? :)
 */

func strStr(haystack string, needle string) int {
    var haystackIdx, needleIdx int

    // 处理异常case
    if len(needle) == 0 {
        return 0
    }

    // 依次遍历haystack，遇到首字母相同的位置S0后，向后和needle的各字母逐个比对，直到完全匹配或者某个字母不匹配
    for haystackIdx = 0; haystackIdx <= len(haystack) - len(needle) && needleIdx < len(needle); haystackIdx++ {
        for needleIdx = 0; needleIdx < len(needle); needleIdx ++ {
            if haystack[haystackIdx+needleIdx] != needle[needleIdx] {
                break
            }
        }
    }

    if needleIdx == len(needle) {
        return haystackIdx - 1
    } else {
        return -1
    }
}