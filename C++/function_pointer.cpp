#include <iostream>
#include <functional>

using namespace std;

template <typename T>
using BinaryFunction = function<T(T, T)>;

template <typename T>
T add(T a, T b)
{
  return a + b;
}

template <typename T>
T subtract(T a, T b)
{
  return a - b;
}

template <typename T>
T multiply(T a, T b)
{
  return a * b;
}

template <typename T>
T divide(T a, T b)
{
  return a / b;
}

template <typename T>
T executeOperation(BinaryFunction<T> operation, T x, T y)
{
  return operation(x, y);
}

int main()
{
  int num1 = 5;
  int num2 = 3;

  // Assign a function pointer
  BinaryFunction<int> op = add<int>;
  cout << "Adding numbers: "<< executeOperation(op, num1, num2) << endl;

  // Reassign to a different function
  op = subtract<int>;
  cout << "Subtracting numbers: "<< executeOperation(op, num1, num2) << endl;

  // Reassign to another function
  op = multiply<int>;
  cout << "Multiplying numbers: "<< executeOperation(op, num1, num2) << endl;

  // Reassign to another function
  op = divide<int>;
  cout << "Dividing numbers: "<< executeOperation(op, num1, num2) << endl;

  // Example with a different data type
  double dNum1 = 5.5;
  double dNum2 = 3.3;
  BinaryFunction<double> opDouble = add<double>;
  cout << "Adding doubles: "<< executeOperation(opDouble, dNum1, dNum2) << endl;

  return 0;
}
