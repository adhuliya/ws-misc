#include <iostream>
#include <string>
#include <unordered_map>

template<class A, class B>
void printMap(std::unordered_map<A,B>& map) {
    std::cout << "THE MAP :" << std::endl;
    for (auto& x : map) {
        std::cout << x.first << " = " << x.second << std::endl;
    }
    std::cout << "=========" << std::endl;
}

int main() {
    std::unordered_map<std::string, int> map;

    map["one"] = 1;
    map["two"] = 2;
    map["three"] = 3;
    map["four"] = 4;

    for (auto& x : map) {
        std::cout << x.first << " = " << x.second << std::endl;
    }

    map.erase("two");

    printMap(map);

    std::cout << "map[\"four\"] = " << map["four"] << std::endl;

    std::string key;
    std::cout << "Enter the key : ";
    getline(std::cin, key);

    if (map.find(key) != map.end()) {
        std::cout << "Found the key, with value = " << map[key] << std::endl;
    }

    std::cout << map.size() << std::endl;
    map.clear();
    std::cout << map.size() << std::endl;

    return 0;
}
