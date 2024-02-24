#include <iostream>
#include <thread>
#include <mutex>

using namespace std;

// Shared resource
class Counter
{
private:
  int _count;

public:
  Counter() : _count(0) {}

  void increment()
  {
    _count++;
  }

  void decrement()
  {
    _count--;
  }

  int getCount()
  {
    return _count;
  }

  void setCountToZero()
  {
    _count = 0;
  }
};

// Problem code without mutex
void problemCode(Counter &counter)
{

  thread t1([&counter]()
            { 
    for (int i = 0; i < 1000000; ++i)
    {
      counter.increment(); 
    } });

  thread t2([&counter]()
            {
    for (int i = 0; i < 1000000; ++i)
    {
      counter.decrement(); 
    } });

  t1.join();
  t2.join();
}

// Solution code with mutex
void solutionCodeMutexLock(Counter &counter)
{
  mutex mtx;

  thread t1([&counter, &mtx]()
            {
      for (int i = 0; i < 1000000; ++i)
      {
        mtx.lock();
        counter.increment();
        mtx.unlock();
      } });

  thread t2([&counter, &mtx]()
            {
      for (int i = 0; i < 1000000; ++i)
      {
        mtx.lock();
        counter.decrement();
        mtx.unlock();
      } });

  t1.join();
  t2.join();
}

int main()
{
  Counter counter;
  problemCode(counter);
  cout << "Problem Count: " << counter.getCount() << endl;

  counter.setCountToZero();

  solutionCodeMutexLock(counter);
  cout << "Solution Count: " << counter.getCount() << endl;

  return 0;
}

/*
Blocking (Lock) vs Non-Blocking (Try Lock) call
Timed Mutex (Try Lock For/Until) - Blocking call for time(t)
Recursive Mutex - Same thread locking the mutex multiple times
  - Unlock should happend in reverse order (ReEntrant Lock)
Lock Guard - Owning the mutex on a scoped basis (No unlock)
Unique Lock - Locking strategies (Defer/Try To/Adopt Lock)
Condition - Inter thread Communication (Notify One/All, Wait)
*/
