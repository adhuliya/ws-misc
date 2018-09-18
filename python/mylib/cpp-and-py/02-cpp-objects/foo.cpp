#include <iostream>

// A C++ program to be used in python

// Compile this file as a shared library to use it in python,
//   g++ -c -fPIC foo.cpp -o foo.o
//   g++ -shared -Wl,-soname,libfoo.so -o libfoo.so  foo.o
// you can also use `clang++` instead.

class Foo{
    public:
        void bar(){
            std::cout << "Hello" << std::endl;
        }
};

// python can only communicate with C type function calls
// so we expose here the "C" type function calls, that
// in turn access the C++ objects and functions.

extern "C" {
    Foo* Foo_new(){ return new Foo(); }
    void Foo_bar(Foo* foo){ foo->bar(); }
    void Foo_del(Foo* foo){ delete foo; }
}
