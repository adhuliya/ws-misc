#include<stdio.h>

int main() {
    int* a;
    int* b;
    int x, y;

    x = 12345;

    a = &x;
    b = &y;

    *b = *a;

    return y;
}
