/*
 * Compile to assembly to see that else part is not eliminated:
 *   clang -O3 optimizeTest.c -S -o optimizeTestClang.s 
 *   gcc -O3 optimizeTest.c -S -o optimizeTestGcc.s
 */

#include<stdio.h>
#include <math.h>

int main() {
  int a, b, c, *u, cond, tmp, input;
  a = 111;
  c = 666;
  u = &a;
  scanf("%d", &input); // special instruction
  input = abs(input);
  cond = input > 0;

  while (cond)  { // `cond` value is undeterministic.
    tmp = *u; // point-of-interest
    b = tmp % c;
    if (b) {
      c = c + 222;
      b = 155;
      u = &b;
    } else {
      u = &c;
      b = 166;
    }
    input = input - 1;
    cond = input > 0;
  }

  printf("%d\n", b);
  return 0;
}

