/*
 * 从尾部开始遍历，加1并逐项进位
 */

func plusOne(digits []int) []int {
    haveCarry := true
    for i := len(digits)-1; i >= 0; i-- {
        if haveCarry {
            digits[i] += 1
        }
        if digits[i] == 10 {
            digits[i] = 0
            haveCarry = true
        } else {
            haveCarry = false
            break
        }
    }

    // 检查首位是否有进位，相应地扩容切片并进位
    if haveCarry {
        tmp := make([]int, 1)
        tmp[0] = 1
        digits = append(tmp, digits...)
    }

    return digits
}