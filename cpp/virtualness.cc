
#include <iostream>
using namespace std;

class A {
    public:
        A() {
        }

        virtual ~A() {
            cout << "Destructed A; x = " << x_ << endl;
        }

    private:
        int x_;
};

class B : public A {
    public:
        B() {
        }

        ~B() {
            cout << "Destructed B;" << endl;
        }
};

void f() {
    A a;
    A aa;
    B b;
}

int main(int argc, char **argv) {
    f();

    cout << "In MAIN:" << endl;
    B *b = new B{};
    delete b;

    A *a = new B{};
    delete a;

    cout << "Main Ends" << endl;

    return 0;
}


