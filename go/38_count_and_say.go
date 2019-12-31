/*
 * 性能区别主要在于字符串拼接方式，此例中strings.Builder优于+
 */

import (
    "strings"
    "strconv"
)

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    var str strings.Builder
    source := countAndSay(n-1)
    count := 1
    say := source[0]
    
    for i := 1; i < len(source); i++ {
        if source[i] != say {
            str.WriteString(strconv.Itoa(count))
            str.WriteByte(say)
            say = source[i]
            count = 1
        } else {
            count++
        }
    }
    str.WriteString(strconv.Itoa(count))
    str.WriteByte(say)
    
    return str.String()
}