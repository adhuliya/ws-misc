
#include <iostream>
#include <vector>


int main(int argc, char **argv) {
  std::vector<int> a;
  a.push_back(10);
  a.push_back(11);
  a.push_back(12);
  a.push_back(13);
  a.push_back(14);

  std::vector<int> b{static_cast<int>(a.size())};

  for (__SIZE_TYPE__ i = 0; i < a.size(); ++i) {
    b[i] = a[i];
  }

  for (__SIZE_TYPE__ i = 0; i < a.size(); ++i) {
    std::cout << b[i] << " ";
  }

  // Min integer value
  std::cout << "\n" << (1 << 31) << "\n"; 
  std::cout << "\n" << (1 << 31) -1 << "\n"; 

  return 0;
}
