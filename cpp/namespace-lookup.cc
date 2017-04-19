
#include <iostream>
using namespace std;

int max = 12;
namespace {
    int max = 10;

    void print()  {
        std::cout << max; // current namespace' max
        std::cout << ::max; // global namespace' max

    }
}

namespace {

    void printt()  {
        std::cout << max; // current namespace' max
        std::cout << ::max; // global namespace' max

    }
}

int main(int argc, char **argv) {
    printt(); // links to the unnamed namespace

    return 0;
}

