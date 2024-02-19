#include <iostream>
#include <climits>
#include <cfloat>

using namespace std;

void printCharRange()
{
  cout << "Size of char: " << sizeof(char) << " bytes" << endl;
  cout << "char: " << CHAR_MIN << " to " << CHAR_MAX << endl;
  cout << "unsigned char: " << 0 << " to " << UCHAR_MAX << endl
       << endl;
}

void printShortRange()
{
  cout << "Size of short: " << sizeof(short) << " bytes" << endl;
  cout << "short: " << SHRT_MIN << " to " << SHRT_MAX << endl;
  cout << "unsigned short: " << 0 << " to " << USHRT_MAX << endl
       << endl;
}

void printIntRange()
{
  cout << "Size of int: " << sizeof(int) << " bytes" << endl;
  cout << "int: " << INT_MIN << " to " << INT_MAX << endl;
  cout << "unsigned int: " << 0 << " to " << UINT_MAX << endl
       << endl;
}

void printLongRange()
{
  cout << "Size of long: " << sizeof(long) << " bytes" << endl;
  cout << "long: " << LONG_MIN << " to " << LONG_MAX << endl;
  cout << "unsigned long: " << 0 << " to " << ULONG_MAX << endl
       << endl;
}

void printLongLongRange()
{
  cout << "Size of long long: " << sizeof(long long) << " bytes" << endl;
  cout << "long long: " << LLONG_MIN << " to " << LLONG_MAX << endl;
  cout << "unsigned long long: " << 0 << " to " << ULLONG_MAX << endl
       << endl;
}

void printFloatRange()
{
  cout << "Size of float: " << sizeof(float) << " bytes" << endl;
  cout << "float: " << FLT_MIN << " to " << FLT_MAX << endl
       << endl;
}

void printDoubleRange()
{
  cout << "Size of double: " << sizeof(double) << " bytes" << endl;
  cout << "double: " << DBL_MIN << " to " << DBL_MAX << endl
       << endl;
}

void printLongDoubleRange()
{
  cout << "Size of long double: " << sizeof(long double) << " bytes" << endl;
  cout << "long double: " << LDBL_MIN << " to " << LDBL_MAX << endl
       << endl;
}

int main()
{
  printCharRange();
  printShortRange();
  printIntRange();
  printLongRange();
  printLongLongRange();
  printFloatRange();
  printDoubleRange();
  printLongDoubleRange();

  return 0;
}
