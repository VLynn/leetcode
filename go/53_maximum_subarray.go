/*
 * 子数组和不大于0，则忽略该子数组，从下一个值开启新的子数组
 * 时间复杂度O(N)
 */

func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    maximum := nums[0]
    for i, sum := 0, 0; i < len(nums); i++ {
        // 累加子数组和，并更新最大值
        sum += nums[i]
        if sum > maximum {
            maximum = sum
        }
        // 子数组和不大于0，则忽略该子数组，从下一个值开启新的子数组
        if sum <= 0 {
            sum = 0
        }
    }

    return maximum
}