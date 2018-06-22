#include<iostream>

#define SIZE 2
int limit;
int divisors[SIZE];

int sum_method_1 () {
  int sum = 0;
  int i,j;

  for (i = 3; i < limit; i++) {
    for (j = 0; j < SIZE; j++) {
      if (i % divisors[j] == 0) {
        sum += i;
        break;
      }
    }
  }

  return sum;
}

int sum_method_2 () {
  int multiple = 1;
  int sum = 0;
  int i;
  for (i = 0; i < SIZE; i++) {
    multiple *= divisors[i];
  }

  for (i = 0; i < SIZE; i++) {
    sum += (divisors[i] * ((limit-1)/divisors[i]) * ((limit-1)/divisors[i] + 1)) >> 1;
  }
  sum -= (multiple * ((limit-1)/multiple) * ((limit-1)/multiple + 1)) >> 1;
  return sum;
}

int main() {
  limit = 1000;
  divisors[0] = 3;
  divisors[1] = 5;

  std::cout << sum_method_1() << std::endl;
  std::cout << sum_method_2() << std::endl;
}
