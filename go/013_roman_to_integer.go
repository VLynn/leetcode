/*
 * 遍历字符串，使用switch将罗马数转成十进制整数，记得处理三个特殊case I,X,C
 */
func romanToInt(s string) int {
    var result int = 0
    var rom2int = map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := 0; i < len(s); i++ {
        if i < len(s)-1 && rom2int[s[i]] < rom2int[s[i+1]] {
            result -= rom2int[s[i]]
        } else {
            result += rom2int[s[i]]
        }
    }
    return result
}