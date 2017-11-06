// stringstreams
#include <iostream>
#include <string>
#include <sstream>
using namespace std;

int main ()
{
  string mystr;
  float price=0;
  int quantity=0;

  cout << "Enter price: ";
  getline (cin,mystr);
  auto ss = stringstream(mystr);
  ss >> price;
  //stringstream(mystr) >> price;

  cout << "Enter quantity: ";
  getline (cin,mystr);
  ss = stringstream(mystr);
  ss >> quantity;
  //stringstream(mystr) >> quantity;

  cout << "Total price: " << price*quantity << endl;
  return 0;
}


