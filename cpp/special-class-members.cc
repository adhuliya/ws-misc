/*
 * [Source](http://www.cplusplus.com/doc/tutorial/classes2/)
 */


#include <iostream>
using namespace std;

class Person {
    public:
        Person() {
            cout << "Person: Default Constructor" << endl;
        }
};

int main(int argc, char **argv) {
    Person p;

    return 0;
}
