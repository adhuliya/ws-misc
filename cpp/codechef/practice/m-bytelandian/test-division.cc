
#include <iostream>
using namespace std;

void printdivision(int i);

int main(int argc, char **argv) {
    for (int i=1; i < 100; i++) {
        std::cout << i << " : ";
        printdivision(i);
        std::cout << std::endl;
    }

    return 0;
}


void printdivision(int i) {
    int a = i >> 1;
    int b = i / 3;
    int c = i >> 2;

    std::cout << "(" << a << ", " << b << ", " << c << ") : " << a + b + c;

    if (i < (a+b+c)) {
        cout << " BIG";
    }
}
