#include <iostream>
#include <thread>
#include <vector>
#include <chrono>

using namespace std;
using namespace std::chrono;

typedef unsigned long long ull;

void oddSum(ull start, ull end)
{
  ull sum = 0;
  for (ull i = start; i <= end; ++i)
  {
    if (i % 2 != 0)
    {
      sum += i;
    }
  }
}

void evenSum(ull start, ull end)
{
  ull sum = 0;
  for (ull i = start; i <= end; ++i)
  {
    if (i % 2 == 0)
    {
      sum += i;
    }
  }
}

int main()
{
  ull start = 1;
  ull end = 1900000000;

  auto startSingle = high_resolution_clock::now();
  oddSum(start, end);
  evenSum(start, end);
  auto endSingle = high_resolution_clock::now();

  auto startMulti = high_resolution_clock::now();
  thread oddThread(oddSum, start, end);
  thread evenThread(evenSum, start, end);
  oddThread.join();
  evenThread.join();
  auto endMulti = high_resolution_clock::now();

  duration<double> singleTime = endSingle - startSingle;
  duration<double> multiTime = endMulti - startMulti;

  cout << "Single-threaded Time: " << singleTime.count() << " seconds" << endl;
  cout << "Multi-threaded Time: " << multiTime.count() << " seconds" << endl;

  return 0;
}

// Single - threaded Time : 6.09766 seconds
// Multi - threaded Time : 3.07827 seconds