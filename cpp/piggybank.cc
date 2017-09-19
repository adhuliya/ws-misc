#include <iostream>

#define ISODD(x) (x & 0x1uL)

int main() {
  int n;
  // bool godsTurn = false;
  unsigned long long p;
  unsigned long long fromGod;

  // No. of test cases
  std::cin >> n;

  while (n--) {
    fromGod = 0;

    // std::cin >> p;
    // if (ISODD(p)) {
    //   godsTurn     = true;
    // } else {
    //   godsTurn     = false;
    // }

    while (p > 0) {
      if (ISODD(p)) {
        p -= 1;
        fromGod += 1;
      } else {
        p >>= 1;
      }
      // if (godsTurn) {
      //   if (ISODD(p)) {
      //     p -= 1;
      //     fromGod += 1;
      //   }
      //   /*else {
      //     // Assume God gave zero
      //   }*/
      //   godsTurn = false;
      // } else {
      //   // should always be even here
      //   if (ISODD(p)) {
      //     std::cout << "ERROR, p is odd after father's turn.\n";
      //   }

      //   p >>= 1;
      //   godsTurn = true;
      // }
    }

    std::cout << fromGod << "\n";
  }

  return 0;
}


