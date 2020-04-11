/*
 * 方法1：二分法
 */

func binarySearch(nums []int, left int, right int, target int) int {
    middle := (left + right) / 2
    
    // 切片中只有一个值时，递归中止
    if left >= right {
        switch {
        case nums[left] < target:
            return left + 1
        case nums[left] > target:
            return left
        case nums[left] == target:
            return left
        default:
            return -1
        }
    }

    // 递归二分
    switch {
    case nums[middle] < target:
        return binarySearch(nums, middle+1, right, target)
    case nums[middle] > target:
        return binarySearch(nums, left, middle-1, target)
    case nums[middle] == target:
        return middle
    default:
        return -1
    }
}

func searchInsert(nums []int, target int) int {
    // 处理空数组case
    if len(nums) == 0 {
        return 0
    }
    
    return binarySearch(nums, 0, len(nums)-1, target)
}


/*
 * 方法二：暴力遍历
 */
func searchInsert(nums []int, target int) int {
    var i int
    for i = 0; i < len(nums); i++ {
        if nums[i] >=  target {
            return i
        }
    }
    return i
}