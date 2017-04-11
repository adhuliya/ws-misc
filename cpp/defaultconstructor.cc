
#include <iostream>

class A {
    public:
        A(int a) {
            std::cout << "A : constructor " << a << std::endl;
        }

        void printmsg() {
            std::cout << "A : printmsg()";
        }
};

class B {
    public:
        A a;

        B() { }
};

int main(int argc, char **argv) {

    B b;

    std::cout << "Hello World";


    return 0;
}
