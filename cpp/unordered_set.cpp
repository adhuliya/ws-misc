#include <iostream>
#include <unordered_set>
#include <string>

template<class A>
void printSet(std::unordered_set<A> &set) {
    for (auto &x : set) {
        std::cout << x << std::endl;
    }
}

int main() {
    using std::string;

    std::unordered_set<string> set;

    std::string str;
    for (int i=0;i < 3; i++) {
        getline(std::cin, str);
        set.insert(str);
    }

    std::cout << "Erase this : ";
    getline(std::cin, str);

    set.erase(str);

    printSet(set);
}
