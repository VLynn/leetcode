/*
 * 给定一个字符串，返回不包含重复字符的最大字串长度
 * 遍历字符串，如果子串中某个字符和遍历到的字符重复了，则将子串的起始点挪到重复字符之后。 
 * 时间复杂度O(n)，空间复杂度O(n)。
 */

func lengthOfLongestSubstring(s string) int {
    // 建立一个hash结构，key是字符串中的某个字符，value是这个字符最近一次出现的位置（靠右的位置）
    byte_hash := make(map[byte]int, len(s))
    // 最大子串的开始位置
    start_pos := 0
    // 最大字串长度，当前子串长度
    longestLen, curLen := 0, 0

    for i := 0; i < len(s); i++ {
        // 如果子串中某个字符和遍历到的字符重复了，则将子串的起始点挪到重复字符之后
        if pos, ok := byte_hash[s[i]]; ok && pos >= start_pos {
            start_pos = pos + 1
            curLen = i - start_pos + 1
        } else {
            curLen++
        }
        // 更新字符最近出现的位置
        byte_hash[s[i]] = i
        // 更新最大子串长度
        if curLen > longestLen {
            longestLen = curLen
        }
    }
    return longestLen
}