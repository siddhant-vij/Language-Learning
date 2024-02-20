#include <iostream>
#define PI 3.14159

using namespace std;

// Base class Shape
class Shape
{
public:
  virtual double getArea() const = 0;
  virtual double getPerimeter() const = 0;
};

// Derived class Circle from Shape
class Circle : public Shape
{
protected:
  double radius;

public:
  Circle(double r) : radius(r) {}

  double getArea() const override
  {
    return PI * radius * radius;
  }

  double getPerimeter() const override
  {
    return 2 * PI * radius;
  }

  void display() const
  {
    cout << "Circle with radius: " << radius << endl;
  }
};

// Base class Colored, a different class hierarchy
class Colored
{
protected:
  string color;

public:
  Colored(string c) : color(c) {}

  string getColor() const
  {
    return color;
  }

  void display() const
  {
    cout << "Color: " << color << endl;
  }
};

// Derived class ColoredCircle from Circle and Colored
class ColoredCircle : public Circle, public Colored
{
public:
  ColoredCircle(double r, string c) : Circle(r), Colored(c) {}

  void display() const
  {
    // Resolve method name conflict by specifying which parent class's method to use
    Circle::display();  // Call display from Circle
    Colored::display(); // Call display from Colored

    // Additional information specific to ColoredCircle can be added here
    cout << "Colored Circle with\n  - Area: " << getArea()
         << "\n  - Perimeter: " << getPerimeter() << endl;
  }
};

// Main function
int main()
{
  ColoredCircle coloredCircle(5, "Red");
  coloredCircle.display();
  return 0;
}
