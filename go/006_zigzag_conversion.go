/*
 * 找规律，(2row-2)个字符是一组
 */

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }

    strSlice := make([]byte, len(s))
    fill_idx := 0
    for i := 0; i < numRows; i++ {
        for index := i; index < len(s); index += (2*numRows-2) {
            strSlice[fill_idx] = s[index]
            fill_idx++
            // 首行和末行之间的行，每组之间，每行需要补一个字符
            if i > 0 && i < numRows-1 && (index + (2*numRows-2) - 2*i) < len(s) {
                strSlice[fill_idx] = s[index + (2*numRows-2) - 2*i]
                fill_idx++
            }
        }
    }
    return string(strSlice)
}