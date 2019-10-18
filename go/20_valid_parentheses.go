/**
 * 采用栈结构，左括号入栈，右括号出栈并配对
 * 空间复杂度O(N)，时间复杂度O(N)
 */

// 弹出slice中最后一个值
func pop(s []byte) (elem byte, result []byte) {
    elem = '0'
    if len(s) > 0 {
        elem = s[len(s)-1]
        s = s[:len(s)-1]
    }
    return elem, s
}

func isValid(s string) bool {
    var elem byte
    stack := make([]byte, 0, len(s))
    var bracketMap = map[byte]byte {
        '(': ')',
        '{': '}',
        '[': ']',
    }
    
    for i := 0; i < len(s); i++ {
        // 左括号入栈，右括号出栈
        if _, ok := bracketMap[s[i]]; ok {
            stack = append(stack, s[i])
        } else {
            elem, stack = pop(stack)
            if elem == '0' || bracketMap[elem] != s[i] {
                return false
            } 
        }
    }
    
    // 循环结束时，栈中应为空
    if len(stack) > 0 {
        return false
    }
    return true
}