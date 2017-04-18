/*
 * Kth largest element from an array.
 */
#include <iostream>
#include <vector>

using namespace std;

template<class T>
void print(const std::vector<T> &vec);
void sort (std::vector<int> &vec);
int kthlargest_sort(std::vector<int> &vec, int k);
int kthlargest_k_insertion_sort(const std::vector<int> &vec, int k);


int main(int argc, char **argv) {
    int kth;
    vector<int> vec{4,5,3,7,8,9,2,10,3,52,645,8,123,567,85};

    print(vec);

    kth = kthlargest_k_insertion_sort(vec, 6);
    cout << "kthlargest : by k element sort : " << kth << endl;

    kth = kthlargest_sort(vec, 6);
    print(vec);
    cout << "kthlargest : bysort : " << kth << endl;

    return 0;
}


template<class T>
void print(const std::vector<T> &vec) {
    std::cout << "Vector : ";
    for(auto i : vec) {
        std::cout << i << " ";
    }

    std::cout << std::endl;
}

void sort (std::vector<int> &vec) {
    int tmp;
    for (int i=0; i < vec.size(); i++) {
        for (int j=0; j < vec.size()-1; j++) {
            if (vec[j] > vec[j+1]) {
                tmp = vec[j];
                vec[j] = vec[j+1];
                vec[j+1] = tmp;
            }
        }
    }
}

int kthlargest_sort(std::vector<int> &vec, int k) {
    sort(vec);

    return vec[k-1];
}

int kthlargest_k_insertion_sort(const std::vector<int> &vec, int k) {
    std::vector<int> v{vec.begin(), vec.begin()+k};

    sort(v);
    print(v);

    v.push_back(0); //AD an extra element

    for (auto i=vec.begin()+k, j=vec.end(); i < j; i++) {
        v[k] = *i; //AD an extra element used here
        for (int n=k-1; n >= 0; n--) {
            if (v[n+1] < v[n]) {
                int tmp = v[n+1];
                v[n+1] = v[n];
                v[n] = tmp;
            }
        }
    }

    return v[k-1];
}


