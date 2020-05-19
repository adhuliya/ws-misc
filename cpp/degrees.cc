#include<iostream>

constexpr long double operator"" _degrees(long double d) {return d * 0.0175;}

int main() {
  std::cout << 180.0_degrees << std::endl;
}
