#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

using namespace std;

class BankAccount
{
private:
  int balance = 0;
  mutex mtx;
  condition_variable cv;

public:
  void add(int value)
  {
    unique_lock<mutex> lock(mtx);
    balance += value;
    cout << "Added " << value << " to the account. New balance: " << balance << endl;
    cv.notify_one();
  }

  void withdraw(int value)
  {
    unique_lock<mutex> lock(mtx);
    cv.wait(lock, [this, value]
            { return balance >= value; });
    // You won't be able to withdraw until the account balance >= withdrawal value. Addition done on separate thread - so this thread is waiting to be notified.
    balance -= value;
    cout << "Withdrawn " << value << " from the account. New balance: " << balance << endl;
  }
};

int main()
{
  BankAccount account;

  thread t1([&account]
            { account.add(50); });

  thread t2([&account]
            { account.withdraw(100); });

  // Since account balance (50) < withdrawal value (100) - deadlock, which can be avoided by below:

  // this_thread::sleep_for(chrono::seconds(3));
  // thread t3([&account]
  //           { account.add(100); });

  // After 3 seconds wait, this thread t3 makes the account balance (150) >= withdrawal value of 100.

  t1.join();
  t2.join();
  // t3.join();

  return 0;
}
