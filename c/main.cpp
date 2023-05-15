#include <stdio.h>
#include <stdlib.h>
// struct
// {
//     int a;
//     char b;
//     int c;
// } t = {10, 'C', 20};

int a = 1;
// 指针变量
int * p = &a;
int main()
{
    // printf("length: %d\n", sizeof(t));
    // printf("&a: %X\n&b: %X\n&c: %X\n", &t.a, &t.b, &t.c);
    // printf("a: %d\nb: %c\nc: %d\n", t.a, t.b, t.c);
    // system("pause");
    printf("a: %d\n", a);
    // 打印指针变量
    printf("p: %X\n", p);
    // 打印指针变量指向的值
    printf("*p: %d\n", *p);
    return 0;
}