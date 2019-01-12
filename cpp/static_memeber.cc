#include<iostream>

class A {
  public:
    static int x;

    void func() {
      x += 1;
      std::cout << x << std::endl;
    }
};

int A::x = 0;

int main() {
  A a = A();
  a.func();
  return 0;
}
