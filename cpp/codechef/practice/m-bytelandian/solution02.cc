
#include <iostream>

using namespace std;

unsigned mem=0;
unsigned max=0;

unsigned dollars(unsigned num, unsigned div, unsigned depth);

int main(int argc, char **argv) {
    unsigned int num;

    while(cin >> num) {
        cout << dollars(num,1,1) << endl;
    }

    cout << "MaxDepth : " << ::max << endl;

    return 0;
}

unsigned dollars(unsigned num, unsigned div, unsigned depth) {
    if (depth > ::max) {
        ::max++;
    }

    num /= div;

    if ((num >= 1 && num <= 11) || (num >= 13 && num <= 15) || num == 17 || num == 19 || num == 23) {
        mem = num/2;
        return num;
    }

    unsigned count=0, sum=0;

    count += dollars(num, 4, depth+1);
    sum += mem + count;

    count += dollars(num, 3, depth+1);
    sum += mem;

    mem = sum;
    count += sum;

    return count;
}


