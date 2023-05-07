#include <stdio.h>

int sum(int a, int b)
{
    return a + b;
}
int main()
{

    // 写一个求和函数并打印结果
    int a = 1, b = 2;
    int c = sum(a, b);
    printf("sum of %d and %d is %d\n", a, b, c);
}
