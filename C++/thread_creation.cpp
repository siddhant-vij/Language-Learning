#include <iostream>
#include <thread>

using namespace std;

void calculateSum(int number)
{
  int sum = 0;
  for (int i = 1; i <= number; ++i)
  {
    sum += i;
  }
  cout << "Sum from 1 to " << number << ": " << sum << endl;
}

// Function pointers
void createThreadUsingFunctionPointer(int number)
{
  thread t1(calculateSum, number);
  t1.join();
}

// Lambda functions
void createThreadUsingLambda(int number)
{
  thread t2([number]()
            { calculateSum(number); });
  t2.join();
}

// Functor (Function Object)
class Functor
{
public:
  void operator()(int number)
  {
    calculateSum(number);
  }
};

void createThreadUsingFunctor(int number)
{
  thread t3((Functor()), number);
  t3.join();
}

// Non-static member function
class NonStaticMember
{
public:
  void sumThread(int number)
  {
    calculateSum(number);
  }
};

void createThreadUsingNonStaticMember(int number)
{
  NonStaticMember nsm;
  thread t4(&NonStaticMember::sumThread, &nsm, number);
  t4.join();
}

// Static member function
class StaticMember
{
public:
  static void staticSumThread(int number)
  {
    calculateSum(number);
  }
};

void createThreadUsingStaticMember(int number)
{
  thread t5(StaticMember::staticSumThread, number);
  t5.join();
}

int main()
{
  int number;
  cout << "Enter a number: ";
  cin >> number;

  createThreadUsingFunctionPointer(number);
  createThreadUsingLambda(number);
  createThreadUsingFunctor(number);
  createThreadUsingNonStaticMember(number);
  createThreadUsingStaticMember(number);

  return 0;
}
