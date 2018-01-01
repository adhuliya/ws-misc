#include<stdio.h>

int addNums(int a, int b) {
    return a + b;
}

// using odd/prime numbers to help locate them easily
int main() {
    int x = addNums(13, 19);

    printf("%d", x);

    return 23;
}
