#include <iostream>
#include <memory>

using namespace std;

class Example
{
public:
  Example()
  {
    cout << "Constructor called." << endl;
  }

  ~Example()
  {
    cout << "Destructor called." << endl;
  }

  void testFunction()
  {
    cout << "Test function called." << endl;
  }
};

// Function demonstrating the use of unique_ptr
void useUniquePointer()
{
  // Creating a unique_ptr to an Example object - Constructor called here
  unique_ptr<Example> uniquePtr = make_unique<Example>();

  // Accessing the managed object (testFunction)
  uniquePtr->testFunction();

  // unique_ptr cannot be copied, only moved
  unique_ptr<Example> movedPtr = move(uniquePtr);

  // After the move, uniquePtr no longer owns the object
  if (!uniquePtr)
  {
    cout << "uniquePtr is now empty." << endl;
  }

  // When movedPtr goes out of scope, the managed object is automatically destroyed - destructor called here
}

// Function demonstrating the use of shared_ptr
void useSharedPointer()
{
  // Creating a shared_ptr to an Example object - Constructor called here
  shared_ptr<Example> sharedPtr1 = make_shared<Example>();

  // sharedPtr2 is now sharing ownership of the same object
  shared_ptr<Example> sharedPtr2 = sharedPtr1;

  // Both pointers can be used to access the testFunction for the same object created with sharedPtr1
  sharedPtr1->testFunction();
  sharedPtr2->testFunction();
  cout << "Reference count: " << sharedPtr1.use_count() << endl;

  // The object will not be destroyed until the last shared_ptr is destroyed - which happens outside this scope & destrucotr is called here
}

// Function demonstrating the use of weak_ptr
void useWeakPointer()
{
  // Creating a shared_ptr like in useSharedPointer
  shared_ptr<Example> sharedPtr = make_shared<Example>();

  // Creating a weak_ptr that observes the shared_ptr
  // The weak_ptr does not own the managed object
  weak_ptr<Example> weakPtr(sharedPtr);

  // The weak_ptr does not affect the reference count of sharedPtr
  cout << "Reference count: " << sharedPtr.use_count() << endl;

  // Locking the weak_ptr to get a shared_ptr that owns the managed object
  shared_ptr<Example> lockedSharedPtr = weakPtr.lock();

  if (lockedSharedPtr)
  {
    // Now we can safely use the object
    lockedSharedPtr->testFunction();
  }
  else
  {
    cout << "The object pointed to by weakPtr no longer exists." << endl;
  }
}

// Case study comparing raw pointers with smart pointers
void rawVsSmartPointers()
{
  // Using raw pointers
  int *rawPtr = new int(100); // Allocate memory for an int
  cout << "Value of rawPtr: " << *rawPtr << endl;
  delete rawPtr; // Manually delete the allocated memory

  // Using smart pointers
  unique_ptr<int> smartPtr(new int(100)); // Allocate memory for an int
  cout << "Value of smartPtr: " << *smartPtr << endl;
  // No need to delete, the memory is automatically managed by the unique_ptr

  // Smart pointers prevent memory leaks and dangling pointers that are common issues with raw pointer usage.
}

int main()
{
  useUniquePointer();
  useSharedPointer();
  useWeakPointer();

  rawVsSmartPointers();

  return 0;
}
