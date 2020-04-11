/*
 * 遇到待删除的元素，从数组最后面挪一个元素过来占坑，并记录已挪过的位置
 */

func removeElement(nums []int, val int) int {
    endIndex := len(nums) - 1
    handleIndex := 0

    for handleIndex <= endIndex {
        if nums[handleIndex] == val {
            nums[handleIndex] = nums[endIndex]
            endIndex--
        } else {
            handleIndex++
        }
    }

    return endIndex + 1
}