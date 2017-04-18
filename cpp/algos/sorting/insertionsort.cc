
#include <iostream>
#include <vector>

using namespace std;


template<class T>
void insertionsort (std::vector<T> &vec) {
    for (int i=1; i < vec.size(); i++) {
        for (int j=0; j < i; j++) {
            if (vec[i-j] < vec[i-j-1]) {
                T tmp = vec[i-j-1];
                vec[i-j-1] = vec[i-j];
                vec[i-j] = tmp;
            }
        }
    }
}


template<class T>
void print (const std::vector<T> &vec) {
    for (auto i : vec) {
        std::cout << i << ", ";
    }
}


int main(int argc, char **argv) {
    std::vector<int> vec {20,30,1,45,7,9,2,90};

    insertionsort(vec);

    print(vec);

    std::cout << std::endl;

    return 0;
}


