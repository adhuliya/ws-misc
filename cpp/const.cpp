
#include <iostream>
using namespace std;

int main(int argc, char **argv) {
    char b = '1';
    char c = '2';
    const char * const a = &b;
    //*a = '3';

    //a = &c;

    return 0;
}
