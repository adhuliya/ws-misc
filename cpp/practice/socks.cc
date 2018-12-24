// John works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
// 
// For example, there are socks with colors . There is one pair of color and one of color . There are three odd socks left, one of each color. The number of pairs is .
// 
// Function Description
// 
// Complete the sockMerchant function in the editor below. It must return an integer representing the number of matching pairs of socks that are available.
// 
// sockMerchant has the following parameter(s):
// 
//     n: the number of socks in the pile
//     ar: the colors of each sock

#include<iostream>
#include<unordered_set>

int main() {
  int n; // size of input
  std::cin >> n;

  int i; // counter
  int color; // sock color
  int pairs = 0; // total sock pairs
  int tmp;

  std::unordered_set<int> us;
  for (i = 0; i < n; ++i) {
    std::cin >> color;
    if (us.find(color) == us.end()) {
      us.insert(color);
    } else {
      pairs += 1;
      us.erase(color);
    }
  }
  std::cout << pairs;
  return 0;
}


