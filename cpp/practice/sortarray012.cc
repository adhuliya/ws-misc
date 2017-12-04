/*
Write a program to sort an array of 0's,1's and 2's in ascending order.

Input:
The first line contains an integer 'T' denoting the total number of test cases. In each test cases, First line is number of elements in array 'N' and second its values.

Output: 
Print the sorted array of 0's, 1's and 2's.

Constraints: 

1 <= T <= 100
1 <= N <= 100
0 <= arr[i] <= 2

Example:

Input :

2
5
0 2 1 2 0
3
0 1 0
 

Output:

0 0 1 2 2
0 0 1
*/


#include <iostream>
#include <vector>
#include <utility>

void sort(std::vector<int>& vec) {
    // basic insertion sort
    int size = vec.size();
    int tmp;

    for (int i=1; i < size; ++i) {
        for (int j=i; j > 0; --j) {
            if (vec[j] < vec[j-1]) {
                //swap j and j-1
                std::swap(vec[j], vec[j-1]);
                //tmp = vec[j];
                //vec[j] = vec[j-1];
                //vec[j-1] = tmp;
            }

        }

    }
}

void print(std::vector<int>& vec) {
    for (auto elem : vec) {
        std::cout << elem << " ";
    }
    std::cout << std::endl;
}

int main() {
    int testCases;
    int vecSize;
    int tmp;
    std::vector<int> vec;

    std::cin >> testCases;

    for (int i=0; i < testCases; ++i) {
        std::cin >> vecSize;
        vec.clear();
        for (int j=0; j < vecSize; ++j) {
            std::cin >> tmp;
            vec.push_back(tmp);
        }
        sort(vec);
        print(vec);
    }

    return 0;
}


