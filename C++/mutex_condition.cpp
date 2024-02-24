#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

using namespace std;

mutex mtx;
condition_variable cv;
bool ready = false;

void mutexLockUnlockDemo()
{
  unique_lock<mutex> lock(mtx);
  cout << "Thread ID: " << this_thread::get_id() << " is inside the critical section" << endl;
  // Simulating some work
  this_thread::sleep_for(chrono::milliseconds(1000));
}

void lockGuardDemo()
{
  lock_guard<mutex> guard(mtx);
  cout << "Thread ID: " << this_thread::get_id() << " is inside the critical section protected by lock guard" << endl;
  // Simulating some work
  this_thread::sleep_for(chrono::milliseconds(1000));
}

void conditionWaitDemo()
{
  unique_lock<mutex> lock(mtx);
  cout << "Thread ID: " << this_thread::get_id() << " is waiting..." << endl;
  cv.wait(lock, []
          { return ready; });
  cout << "Thread ID: " << this_thread::get_id() << " is notified and continuing" << endl;
}

void conditionNotifyDemo()
{
  this_thread::sleep_for(chrono::milliseconds(2000));
  {
    lock_guard<mutex> guard(mtx);
    ready = true;
    cout << "Thread ID: " << this_thread::get_id() << " Notifying all threads..." << endl;
  }
  cv.notify_all();
}

int main()
{
  cout << "Starting the program on main thread with ID: " << this_thread::get_id() << endl;

  thread t2(mutexLockUnlockDemo);
  thread t3(lockGuardDemo);
  thread t4(conditionWaitDemo);
  thread t5(conditionNotifyDemo);

  t2.join();
  t3.join();
  t4.join();
  t5.join();

  return 0;
}
