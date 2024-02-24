#include <iostream>
#include <future>
#include <thread>
#include <chrono>

typedef unsigned long long ull;

using namespace std;

ull oddSumAndSetPromise(ull start, ull end)
{
  ull sum = 0;
  for (ull i = start; i <= end; i++)
  {
    if (i % 2 != 0)
      sum += i;
  }
  return sum;
}

int main()
{
  future<ull> oddSumFuture = async(launch::deferred, oddSumAndSetPromise, 1, 1900000000); // Deferred

  ull sumResult = oddSumFuture.get(); // Blocking call

  cout << "Sum of odd numbers from 1 to 1900000000: " << sumResult << endl;

  return 0;
}
