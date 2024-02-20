#include <iostream>

using namespace std;

class Complex
{
private:
  double _real;
  double _imaginary;

public:
  Complex(double real, double imaginary)
      : _real(real), _imaginary(imaginary) {}

  Complex operator+(const Complex &rhs) const
  {
    return Complex(this->_real + rhs._real, this->_imaginary + rhs._imaginary);
  }

  Complex operator-(const Complex &rhs) const
  {
    return Complex(this->_real - rhs._real, this->_imaginary - rhs._imaginary);
  }

  Complex operator*(const Complex &rhs) const
  {
    return Complex(this->_real * rhs._real - this->_imaginary * rhs._imaginary, this->_real * rhs._imaginary + this->_imaginary * rhs._real);
  }

  void print() const
  {
    cout << "(" << this->_real << " + " << this->_imaginary << "i)" << endl;
  }
};

int main()
{
  double real1, imaginary1, real2, imaginary2;
  cout << "Enter real and imaginary parts of the first complex number: ";
  cin >> real1 >> imaginary1;
  cout << "Enter real and imaginary parts of the second complex number: ";
  cin >> real2 >> imaginary2;

  Complex complex1(real1, imaginary1);
  Complex complex2(real2, imaginary2);

  Complex sum = complex1 + complex2;
  Complex difference = complex1 - complex2;
  Complex product = complex1 * complex2;

  cout << "The sum of the complex numbers is: ";
  sum.print();
  cout << "The difference of the complex numbers is: ";
  difference.print();
  cout << "The product of the complex numbers is: ";
  product.print();

  return 0;
}
