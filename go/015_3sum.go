/*
 * 拆成正负两个数组，两负一正 or 一负两正 or 三零 （零分到正组）
 * 时间复杂度O(N3)，但常数系数比较小
 * Runtime: 2472 ms, faster than 5.00% of Go online submissions for 3Sum.
 * Memory Usage: 7.5 MB, less than 100.00% of Go online submissions for 3Sum.
 */

func threeSum(nums []int) [][]int {
    positives := make([]int, 0, len(nums))
    negatives := make([]int, 0, len(nums))
    res := make([][]int, 0, len(nums)/3)

    // 把原数组分为正、负两个数组
    for _, v := range nums {
        if v >= 0 {
            positives = append(positives, v)
        } else {
            negatives = append(negatives, v)
        }
    }

    // 两负一正
    if len(negatives) >= 2 && len(positives) >= 1 {
        for i := 0; i < len(negatives) - 1; i++ {
            for j := i + 1; j < len(negatives); j++ {
                for k := 0; k < len(positives); k++ {
                    if negatives[i] + negatives[j] + positives[k] == 0 {
                        // 元组按从小到大的顺序写入切片，方便后续去重
                        if negatives[i] > negatives[j] {
                            res = append(res, []int{negatives[j], negatives[i], positives[k]})
                        } else {
                            res = append(res, []int{negatives[i], negatives[j], positives[k]})
                        }
                    }
                }
            }
        }
    }
    
    // 两正一负
    if len(positives) >= 2 && len(negatives) >= 1 {
        for i := 0; i < len(positives) - 1; i++ {
            for j := i + 1; j < len(positives); j++ {
                for k := 0; k < len(negatives); k++ {
                    if positives[i] + positives[j] + negatives[k] == 0 {
                        if positives[i] > positives[j] {
                            res = append(res, []int{negatives[k], positives[j], positives[i]})
                        } else {
                            res = append(res, []int{negatives[k], positives[i], positives[j]})
                        }
                    }
                }
            }
        }
    }

    // 三零
    zero_times := 0
    for _, v := range positives {
        if v == 0 {
            zero_times++
        }
        if zero_times >= 3 {
            res = append(res, []int{0, 0, 0})
        }
    }

    // 去重
    undup_res := make([][]int, 0, len(res))
    LOOP:
    for i := 0; i < len(res); i++ {
        for j := i + 1; j < len(res); j++ {
            if res[i][0] == res[j][0] && res[i][1] == res[j][1] && res[i][2] == res[j][2] {
                continue LOOP
            }
        }
        undup_res = append(undup_res, res[i])
    }

    return undup_res
}