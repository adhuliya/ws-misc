
#include <iostream>

using namespace std;

unsigned dollars(unsigned n, unsigned onefourth);

int main(int argc, char **argv) {
    unsigned int n;

    while(cin >> n) {
        cout << dollars(n,0) << endl;
    }

    return 0;
}

unsigned dollars(unsigned n, unsigned onefourth) {

    if ((n > 0 && n < 10) || (n >= 13 && n <= 15) || n == 17 || n == 19 || n == 23) return n;

    if (onefourth) {
        return onefourth + dollars(n>>2, 0) + dollars(n/3,0);
    }

    unsigned onefourth_ = dollars(n>>2,0);
    return onefourth_ + dollars(n>>1, onefourth_) + dollars(n/3,0);

}
