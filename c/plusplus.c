#include <stdio.h>

int main() {
  int x;

  // CASE 1
  x = 1;
  x = x++;
  printf("x = 1; x = x++; then x = %d\n", x);

  // CASE 2
  x = 1;
  x = ++x;
  printf("x = 1; x = ++x; then x = %d\n", x);

  // CASE 3
  int z[2], *a;
  z[0] = 10; z[1] = 20;
  a = z;
  *a = *(++a);
  printf("z[0] = 10; z[1] = 20; a = z; *a = *(++a); then z[0] = %d, z[1] = %d\n", z[0], z[1]);

  // CASE 4
  int u[2], *b;
  u[0] = 10; u[1] = 20;
  b = u;
  *b = *b + *(++b);
  printf("u[0] = 10; u[1] = 20; b = u; *b = *b + *(++b); then u[0] = %d, u[1] = %d\n", u[0], u[1]);

  return 0;
}
