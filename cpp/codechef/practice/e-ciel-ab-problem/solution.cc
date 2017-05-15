
#include <iostream>
using namespace std;

int main(int argc, char **argv) {
    int a, b;

    cin >> a >> b;

    int c = a - b;
    int d = c % 10;

    if (d == 1) {
        d = 2;
    } else {
        d = 1;
    }

    cout << (c/10)*10 + d;

    return 0;
}
