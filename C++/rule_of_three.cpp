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

  // Copy assignment operator
  Rectangle &operator=(const Rectangle &other)
  {
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

  // Destructor
  ~Rectangle()
  {
    delete[] arr;
  }

  void print()
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
  r1.print();

  Rectangle r2(r1);
  r2.print();

  Rectangle r3;
  r3 = r1;
  r3.print();

  return 0;
}
