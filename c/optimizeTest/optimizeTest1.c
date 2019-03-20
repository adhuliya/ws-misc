#include<stdio.h>

int main() {
  int a, b, c, *u, cond, tmp, input;
  a = 111;
  u = &a;

  scanf("%d", &cond); // special instruction
  while (cond)  { // `cond` value is undeterministic (statically)
    tmp = *u; // point-of-interest
    if (tmp) {
      b = 155;
      u = &b;
    } else {
      b = 166;
    }
    cond = cond - 1;
  }

  printf("%d\n", b);
  return 0;
}
