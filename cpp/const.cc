
#include <iostream>
using namespace std;

int main(int argc, char **argv) {
    int a = 10;
    int b = 20;
    const int *x = &a;
    int*const y = &a;
    const int * const z = &a;

    //*x = 20;
    *y = 30;

    x = &b;
    //AD y = &b;

    cout << *x << ", " << *y << ", " << *z;

    return 0;
}
