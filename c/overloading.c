/*
 * This program doesn't compile, demonstrating that
 * C doesn't support overloading.
 */

#include <stdio.h>

void print(int x) {
  printf("INT: %d\n", x);
}

void print(float f) {
  printf("FLOAT: %f\n", f);
}

void main() {
  int x = 1;
  float f = 10.5;

  print(x);
  print(f);
}
