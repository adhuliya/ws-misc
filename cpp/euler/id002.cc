// Sum of even fibonacci terms <= 4000000
#include <iostream>

#define MAX 4000000

int main(int argc, char **argv) {
  int sum = 2;
  int a = 1;
  int b = 2;
  int tmp;
  
  while(1) {
    tmp = a + b;
    a = b + tmp;
    b = a + tmp;
    if (b <= MAX) {
      sum += b;
    } else {
      break;
    }
  }

  std::cout << sum << std::endl;
  return 0;
}


