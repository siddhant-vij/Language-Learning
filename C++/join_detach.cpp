#include <iostream>
#include <thread>
#include <vector>

using namespace std;

// Function to sum elements of an array
void sumArrayElements(int *arr, int size, int &result)
{
  result = 0;
  for (int i = 0; i < size; ++i)
  {
    result += arr[i];
  }
}

int main()
{
  int arr[] = {1, 2, 3, 4, 5};
  int size = sizeof(arr) / sizeof(arr[0]);
  int result;

  // Creating a thread to sum the array elements
  thread t(sumArrayElements, arr, size, ref(result));

  // ------------- Testing code -------------
  // t.join();

  // Check if the thread is joinable
  if (t.joinable())
  {
    cout << "Thread is joinable" << endl;

    // Wait for the thread to finish
    t.join();

    cout << "Sum of array elements: " << result << endl;
  }
  else
  {
    cout << "Thread is not joinable" << endl;

    // Detach the thread
    // Note: Check if it's joinable before detaching
    if (t.joinable())
    {
      t.detach();
    }
  }

  return 0;
}

// Double join & detach will crash the program