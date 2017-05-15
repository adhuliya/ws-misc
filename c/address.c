
#include <stdio.h>
int main() {
    int a = 15, *x = &a;
    printf("%p\n%p\n%p\n%p\n", x, *(&x), &a, &(*x));
}
