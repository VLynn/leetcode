/*
 * 遍历字符串，使用switch将罗马数转成十进制整数，记得处理三个特殊case I,X,C
 */
func romanToInt(s string) int {
    var result int = 0;
    for i := 0; i < len(s); i++ {
        switch s[i] {
        case 'I':
            if (i < len(s)-1) && (s[i+1] == 'V' || s[i+1] == 'X') {
                result -= 1
            } else {
                result += 1
            }
        case 'V':
            result += 5
        case 'X':
            if (i < len(s)-1) && (s[i+1] == 'L' || s[i+1] == 'C') {
                result -= 10
            } else {
                result += 10
            }
        case 'L':
            result += 50
        case 'C':
            if (i < len(s)-1) && (s[i+1] == 'D' || s[i+1] == 'M') {
                result -= 100
            } else {
                result += 100
            }
        case 'D':
            result += 500
        case 'M':
            result += 1000
        }
    }
    return result
}