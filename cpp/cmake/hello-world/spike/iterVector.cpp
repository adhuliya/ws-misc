#include <vector>
#include <iostream>

int main() {
  std::vector<int> intVec;
  intVec.push_back(10);
  intVec.push_back(20);

  for (auto val : intVec) {
    std::cout << val << std::endl;
  }

  return 0;
}
