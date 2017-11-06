
#include <iostream>

using namespace std;

int& select(int& a, int& b) {
    if (a > b) {
        return a;
    } else {
        return b;
    }
}

int main(int argc, char **argv) {
    int a = 10, b=20;
    int c = 20, d=10;
    
    select(a,b) = 100;
    select(c,d) = 100;

    cout << a << " " << b << endl;
    cout << c << " " << d << endl;

    return 0;
}
