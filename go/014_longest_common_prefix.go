// 逐个字符遍历，直到某个字符相异，或者某个字符串已遍历完
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    index := 0
    outer:
    for {
        if len(strs[0]) < index + 1 {
            break outer
        } 
        ch := strs[0][index]
        
        for _, str := range strs {
            if len(str) < index + 1 || str[index] != ch {
                break outer
            }
        }
        index++
    }
    return strs[0][:index]
}