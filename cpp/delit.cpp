
#include <iostream>
#include <vector>

void f() {}

int main() {
  std::cout << (f(), "hello") << "\n";
  return 0;
}
