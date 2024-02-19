#include <iostream>

using namespace std;

template <typename T>
void printAll(T t)
{
  cout << t << endl;
}

template <typename T, typename... Args>
void printAll(T t, Args... args)
{
  cout << t << endl;
  printAll(args...);
}

int main()
{
  printAll(1, 2.5, "three", '4', true, 2.9f, -4, "five", false);

  return 0;
}
