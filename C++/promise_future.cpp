#include <iostream>
#include <future>
#include <thread>
#include <chrono>

typedef unsigned long long ull;

using namespace std;

void oddSumAndSetPromise(ull start, ull end, promise<ull> &oddSumPromise)
{
  ull sum = 0;
  for (ull i = start; i <= end; i++)
  {
    if (i % 2 != 0)
      sum += i;
  }

  // Setting the value of promise
  oddSumPromise.set_value(sum);
}

int main()
{
  // Creating a promise to store the sum
  promise<ull> oddSumPromise;

  // Getting the future associated with the promise
  future<ull> oddSumFuture = oddSumPromise.get_future();

  // Creating a thread to calculate sum and set the promise
  thread oddSumThread(oddSumAndSetPromise, 1, 1900000000, ref(oddSumPromise));

  // Getting the value from the future
  ull sumResult = oddSumFuture.get();

  // Joining the thread
  oddSumThread.join();

  cout << "Sum of odd numbers from 1 to 1900000000: " << sumResult << endl;

  return 0;
}
