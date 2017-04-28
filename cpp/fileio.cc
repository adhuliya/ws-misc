
#include <iostream>
#include <fstream>
using namespace std;

int main(int argc, char **argv) {

    ofstream out;
    out.open("test.txt", ios::app);

    out << "hello world\n";

    return 0;
}
