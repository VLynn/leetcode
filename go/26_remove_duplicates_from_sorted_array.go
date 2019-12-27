/*
 * 遍历数组中的所有元素，一旦出现比上一个插入元素大的元素，就将它插入到正确的位置
 */
 
func removeDuplicates(nums []int) int {
    // 待处理的下一个元素索引
    handleIndex := 1
    // 待插入的下一个位置
    insertIndex := 1
    
    for handleIndex < len(nums) {
        if nums[handleIndex] > nums[insertIndex - 1] {
            nums[insertIndex] = nums[handleIndex]
            insertIndex++
        }
        handleIndex++
    }
    
    return insertIndex
}