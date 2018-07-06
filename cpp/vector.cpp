#include<iostream>
#include<vector>

using namespace std;

int main() {
    vector<int> vec;

    vec.push_back(10);
    vec.push_back(20);
    vec.push_back(21);
    vec.push_back(22);

    vec.pop_back();

    vec[2] = 34;

    for(int i=0, j=vec.size(); i < j; i++) {
        cout << vec[i] << endl;
    }

    cout << endl;

    for(auto i : vec) {
        cout << i << endl;
    }

    cout << endl;

    for(auto i = vec.begin(), j = vec.end(); i < j; i++) {
        cout << *i << endl;
    }

    cout << "Size(Before) : " << vec.size() << endl;
    vec.clear();
    cout << "Size(After)  : " << vec.size() << endl;

    return 0;
}
