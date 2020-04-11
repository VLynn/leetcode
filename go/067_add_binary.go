/*
 * 选取较长的那个字符串开始遍历，进位加1
 * 逐位相加时，注意字节值和字符值的区别
 */

func addBinary(a string, b string) string {
    // 字符串转字节数组，方便操作
    var long, short []byte
    if len(a) > len(b) {
        long, short = []byte(a), []byte(b)
    } else {
        long, short = []byte(b), []byte(a)
    }

    haveCarry := false
    for i := 1; i <= len(long); i++ {
        // 短字符串已累加完，且没有进位，结束循环
        if i > len(short) && !haveCarry {
            return string(long)
        }
        // 累加短字符串
        if i <= len(short) {
            long[len(long)-i] += (short[len(short)-i] - '0')
        }
        // 进位加1
        if haveCarry {
            long[len(long)-i] += 1
        }
        // 检查是否进位
        if long[len(long)-i] >= '2' {
            haveCarry = true
            long[len(long)-i] -= 2
        } else {
            haveCarry = false
        }
    }

    if haveCarry {
        return "1" + string(long)
    }
    return string(long)
}