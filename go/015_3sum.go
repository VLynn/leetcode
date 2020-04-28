Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

/*
 * 拆成正负两个数组，两负一正 or 一负两正 or 三零 （零分到正组）
 */

func threeSum(nums []int) [][]int {
    positives := make([]int, 0, len(nums))
    negatives := make([]int, 0, len(nums))
    res := make([][]int, 0, len/3)

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
                        res = append(res, []int{negatives[i], negatives[j], positives[k]})
                    }
                }
            }
        }
    }
    
    // 两正一负
    if len(positives) >= 2 && len(negatives) >= 1 {
        for 
    }

    // 三零
    if len(positives) >= 3 {

    }
}