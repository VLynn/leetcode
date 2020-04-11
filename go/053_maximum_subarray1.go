/*
 * 子数组和不大于0，则忽略该子数组，从下一个值开启新的子数组
 * 时间复杂度O(N)
 * 
 * 假设Ai, A(i+1), ... Aj是从Ai起始的首个负值子序列， O(N)的算法证明包含两部分：
 * 1. 在i-j范围内，包含Aj的所有后缀子序列都不能作为最优子序列的前缀（包含Aj的所有后缀子序列都是负值）
 * 2. 包含Ai的最大子序列和就是i-j范围内的最大子序列和（包含Ai的所有前缀子序列和都是正值）
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