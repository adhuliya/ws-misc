
#include <iostream>
#include <unordered_map>

int main() {
    std::unordered_map<std::string, int> numberMap;

    // Adding values: simple way
    numberMap["one"] = 1;
    numberMap["two"] = 8; // deliberately wrong
    numberMap["three"] = 3;
    numberMap["four"] = 4;

    // Adding values: complex way
    numberMap.insert(std::make_pair("five", 5));
    // numberMap.insert(std::make_pair("six", 6));

    // Modify values
    numberMap["two"] = 2;

    // Delete key
    numberMap.erase("six");


    // Iterate throught the map
    for (auto it = numberMap.begin(); it != numberMap.end(); ++it) {
        // two ways to access values: it->second and numberMap[it->first]
        std::cout << it->first << ", " << it->second << ":" << numberMap[it->first] << std::endl;
    }

    // Find the element.
    if(numberMap.find("six") != numberMap.end()) {
        std::cout << "Six exists.";
    } else {
        std::cout << "Six does not exist.";
    }

    // Remove an element
}


