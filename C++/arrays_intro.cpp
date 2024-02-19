#include <iostream>

using namespace std;

void printArray(int *arr, int size)
{
  cout << "[";
  for (int i = 0; i < size; ++i)
  {
    cout << arr[i]; // Using array indexes
    if (i < size - 1)
      cout << ", ";
  }
  cout << "]" << endl;
}

void getArrayValues(int *arr, int size)
{
  cout << "[";
  for (int i = 0; i < size; ++i)
  {
    cout << *(arr + i); // Pointer Arithmetic
    if (i < size - 1)
      cout << ", ";
  }
  cout << "]" << endl;
}

void updateArray(int *arr, int size, int index, int newValue)
{
  if (index >= 0 && index < size)
  {
    // arr[index] = newValue; // Using array indexes
    *(arr + index) = newValue; // Pointer Arithmetic
  }
}

void modifyArray(int (&arr)[], int size, int newValue)
{
  // Update elements to newValue using reference to array
  for (int i = 0; i < size; ++i)
  {
    arr[i] = newValue;
  }
}

int main()
{
  int size;
  cin >> size;
  int arr[size];
  for (int i = 0; i < size; i++)
  {
    cin >> arr[i];
  }
  getArrayValues(arr, size);
  printArray(arr, size);

  int indexToUpdate, newValue;
  cout << "Enter index to update: ";
  cin >> indexToUpdate;
  cout << "Enter new value: ";
  cin >> newValue;
  updateArray(arr, size, indexToUpdate, newValue);
  printArray(arr, size);

  int newUniformValue;
  cout << "Enter new value to set all elements: ";
  cin >> newUniformValue;
  modifyArray(arr, size, newUniformValue);
  printArray(arr, size);
  return 0;
}