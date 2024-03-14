#include <iostream>

using namespace std;

class Rectangle
{
private:
  int *arr;
  int size;

public:
  // Default constructor
  Rectangle()
  {
    size = 0;
    arr = new int[size];
  }

  // Parameterized constructor
  Rectangle(int s)
  {
    size = s;
    arr = new int[size];

    for (int i = 0; i < size; i++)
    {
      arr[i] = i * i;
    }
  }

  // Copy constructor
  Rectangle(const Rectangle &other)
  {
    size = other.size;
    arr = new int[size];

    for (int i = 0; i < size; i++)
    {
      arr[i] = other.arr[i];
    }
  }

  // Move constructor
  Rectangle(Rectangle &&other) noexcept
  {
    arr = other.arr;
    size = other.size;

    other.arr = nullptr;
    other.size = 0;
  }

  // Copy assignment operator
  Rectangle &operator=(const Rectangle &other)
  {
    // Self-assignment guard
    if (this != &other)
    {
      delete[] arr;
      size = other.size;
      arr = new int[size];

      for (int i = 0; i < size; i++)
      {
        arr[i] = other.arr[i];
      }
    }

    return *this;
  }

  // Move assignment operator
  Rectangle &operator=(Rectangle &&other) noexcept
  {
    // Self-assignment guard
    if (this != &other)
    {
      delete[] arr;
      arr = other.arr;
      size = other.size;

      other.arr = nullptr;
      other.size = 0;
    }

    return *this;
  }

  // Destructor
  ~Rectangle()
  {
    delete[] arr;
  }

  void printValues()
  {
    for (int i = 0; i < size; i++)
    {
      cout << arr[i] << " ";
    }
    cout << endl;
  }
};

int main()
{
  Rectangle r1(5);
  r1.printValues();

  Rectangle r2(10);
  r2.printValues();

  Rectangle r3 = r1;
  r3.printValues();

  Rectangle r4(15);
  r4 = r2;
  r4.printValues();

  Rectangle r5(20);
  r5 = move(r3);
  r5.printValues();

  return 0;
}
