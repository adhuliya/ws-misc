#include <iostream>

class A {
  A() {}
  public:
  int x;

  friend class B;
};

class B {
  public:
    static A getA() {
      A a{};
      return a;
    }
};

int main() {
  A a = B::getA(); 
  std::cout << "Hello\n";
  return 0;
}

