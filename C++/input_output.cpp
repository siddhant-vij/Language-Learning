#include <iostream>
#include <string>

using namespace std;

int main() {
  string name;
  cout << "Enter your name: ";
  getline(cin, name);
  cout << "Hello, " << name << ". Hope your day is going great!" << endl;
  return 0;
}