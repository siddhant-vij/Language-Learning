#include <iostream>

using namespace std;

// Iterative approach to calculate factorial
unsigned long long factorialIterative(unsigned int n)
{
  unsigned long long factorial = 1;
  for (unsigned int i = 1; i <= n; ++i)
  {
    factorial *= i;
  }
  return factorial;
}

// Recursive approach to calculate factorial
unsigned long long factorialRecursive(unsigned int n)
{
  if (n == 0)
    return 1;
  return n * factorialRecursive(n - 1);
}

int main()
{
  int number;
  cout << "Enter a number to find its factorial: ";
  cin >> number;

  if (number < 0)
  {
    cout << "Factorial is not defined for negative numbers." << endl;
    return 0;
  }

  unsigned long long iterativeResult = factorialIterative(number);
  cout << "Iterative: " << number << "! = " << iterativeResult << endl;

  unsigned long long recursiveResult = factorialRecursive(number);
  cout << "Recursive: " << number << "! = " << recursiveResult << endl;

  return 0;
}
