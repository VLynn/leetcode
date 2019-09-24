// 反转整数
int reverseInt(int x) {
    int result = 0;
    while (fabs(x) > 9) {
        result += x % 10;
        // 处理越界
        if (result > pow(2, 31) / 10) {
            return 0;
        }
        result *= 10;
        x /= 10;
    }
    return result + x % 10;
}

bool isPalindrome(int x){
    if (x < 0) {
        return false;
    }
    if (reverseInt(x) == x) {
        return true;
    }
    return false;
}