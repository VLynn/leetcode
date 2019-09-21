/**
 * 暴力破解法
 * 时间复杂度O(N2)，空间复杂度O(1)
 *
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    int* returns = malloc(sizeof(int)*2);
    int i, j;
    
    for (i = 0; i < numsSize - 1; i++) {
        for (j = i + 1; j < numsSize; j++) {
            if (nums[j] == target - nums[i]) {
                returns[0] = i;
                returns[1] = j;
                *returnSize = 2;
                return returns;
            }
        }
    }
    
    *returnSize = 0;
    return NULL;
}