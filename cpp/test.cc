#include <iostream>

#define ISODD(x) (x & 0x1uL)

int main() {
  int n;
  unsigned long long p;
  unsigned long long fromGod;

  // No. of test cases
  std::cin >> n;

  while (n--) {
    fromGod = 0;

    std::cin >> p;
    while (p > 0) {
      if (ISODD(p)) {
        fromGod += 1;
      }
      p >>= 1;
    }

    std::cout << fromGod << "\n";
  }

  return 0;
}
