#include <iostream>

using namespace std;

double calculator(double num1, char op, double num2)
{
  if (op == '+')
  {
    return num1 + num2;
  }
  else if (op == '-')
  {
    return num1 - num2;
  }
  else if (op == '*')
  {
    return num1 * num2;
  }
  else if (op == '/')
  {
    return num1 / num2;
  }
  else
  {
    return 0;
  }
}

int main()
{
  double num1, num2;
  char op;
  cout << "Enter two numbers: ";
  cin >> num1 >> num2;
  cout << "Enter an operator (+, -, *, /): ";
  cin >> op;
  cout << num1 << " " << op << " " << num2 << " = " << calculator(num1, op, num2) << endl;
  return 0;
}