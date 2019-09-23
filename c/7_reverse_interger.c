/*
 * We want to repeatedly "pop" the last digit off of xx and "push" it to the back of the rev.
 * In the end, rev will be the reverse of the xx.
 */
int reverse(int x) {    
    int result = 0;
    while (fabs(x) > 9) {
        result += x % 10;
        // 边界检查
        if (fabs(result) > (pow(2, 31) / 10)) {
            return 0;
        }
        result *= 10;
        x = x / 10;
    }
    return result + x % 10;
}