#include <iostream>
#include <thread>
#include <mutex>

using namespace std;

mutex mtx1, mtx2;

void thread1DL()
{
  // Lock mutex1
  unique_lock<mutex> lock1(mtx1);
  cout << "Thread 1 locked mutex1" << endl;

  // Simulate some work
  this_thread::sleep_for(chrono::milliseconds(1000));

  // Try to lock mutex2
  unique_lock<mutex> lock2(mtx2);
  cout << "Thread 1 locked mutex2" << endl;
}

void thread2DL()
{
  // Lock mutex2
  unique_lock<mutex> lock2(mtx2);
  cout << "Thread 2 locked mutex2" << endl;

  // Simulate some work
  this_thread::sleep_for(chrono::milliseconds(1000));

  // Try to lock mutex1
  unique_lock<mutex> lock1(mtx1);
  cout << "Thread 2 locked mutex1" << endl;
}

void thread1Solution()
{
  // Lock mutex1
  unique_lock<mutex> lock1(mtx1);
  cout << "Thread 1 locked mutex1" << endl;

  // Simulate some work
  this_thread::sleep_for(chrono::milliseconds(1000));

  cout << "Thread 1 locked mutex2" << endl;
  mtx2.unlock();
}

void thread2Solution()
{
  // Lock mutex2
  unique_lock<mutex> lock2(mtx2);
  cout << "Thread 2 locked mutex2" << endl;

  // Simulate some work
  this_thread::sleep_for(chrono::milliseconds(1000));

  // Try to lock mutex1
  if (mtx1.try_lock())
  {
    cout << "Thread 2 locked mutex1" << endl;
    mtx1.unlock();
  }
  else
  {
    cout << "Thread 2 couldn't lock mutex1 - avoiding deadlock" << endl;
  }
}

int main()
{
  // thread t1(thread1DL);
  // thread t2(thread2DL);

  // t1.join();
  // t2.join();

  // ------------------------------------------

  thread t3(thread1Solution);
  thread t4(thread2Solution);

  t3.join();
  t4.join();

  return 0;
}
