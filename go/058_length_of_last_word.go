/*
 * 从字符串尾部开始遍历
 */

func lengthOfLastWord(s string) int {
    length := 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ' ' && length > 0 {
            break
        }
        if s[i] != ' ' {
            length++
        }
    }
    return length
}