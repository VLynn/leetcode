/*
 * 暴力破解法
 * 时间复杂度O(N2)，空间复杂度O(1)
 */
 
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}