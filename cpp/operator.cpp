#include <iostream>
#include <vector>


class MyArray {
    public:
        MyArray() {
            for (int i=0; i < 10; i++) {
                vec.push_back(i);
            }
        }

        int& operator[](const int i) {
            return vec[i];
        }

        int size() {
            return vec.size();
        }

    private:
        std::vector<int> vec;
};


int main() {
    MyArray arr;

    arr[2] = 22;

    for(int i = 0; i < arr.size(); i++) {
        std::cout << arr[i] << std::endl;
    }
}


