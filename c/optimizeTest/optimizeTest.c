#include<stdio.h>

int main() {
  int a, b, c, *u, cond, tmp, input;
  a = 111;
  u = &a;
  scanf("%d", &input); // special instruction

  cond = input > 0;
  while (cond)  { // `cond` value is undeterministic (statically)
    tmp = *u; // point-of-interest
    b = tmp % 2;
    if (b) {
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
