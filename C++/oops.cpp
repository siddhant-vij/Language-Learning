#include <iostream>
#include <string>

using namespace std;

// Forward declaration of class Person
class Person;

// Friend class declaration
class OccupationViewer
{
public:
  // Friend method that can access private and protected members of Person
  static void viewOccupation(const Person &p);
};

// Base class Person
class Person
{
private:
  // Private data members for Encapsulation
  string _name;
  int _age;

public:
  // Public constructor - Abstraction
  Person(string name, int age)
      : _name(name), _age(age) {}

  // Virtual method - Polymorphism (method overriding)
  virtual void displayOccupation() const
  {
    cout << _name << " is a person." << endl;
  }

  // Getters
  string getName() const { return _name; }
  int getAge() const { return _age; }

  // Declare OccupationViewer as a friend class
  friend class OccupationViewer;
};

// Define the friend method that can access private members of Person
void OccupationViewer::viewOccupation(const Person &p)
{
  p.displayOccupation();
}

// Employee derived from Person - Inheritance
class Employee : public Person
{
protected:
  // Protected data member accessible in derived classes
  string company;

public:
  // Constructor using base class constructor
  Employee(string name, int age, string company)
      : Person(name, age), company(company) {}

  // Method overriding
  void displayOccupation() const override
  {
    cout << getName() << " is an employee at " << company << "." << endl;
  }
};

// Further derived class SoftwareEngineer from Employee
class SoftwareEngineer : public Employee
{
public:
  // Constructor using base class constructor
  SoftwareEngineer(string name, int age, string company)
      : Employee(name, age, company) {}

  // Method overriding
  void displayOccupation() const override
  {
    cout << getName() << " is a software engineer at " << company << "." << endl;
  }
};

// Example of Heirarchical Inheritance
class FAANGEngineer : public SoftwareEngineer
{
public:
  // Constructor using base class constructor
  FAANGEngineer(string name, int age, string company) : SoftwareEngineer(name, age, company) {}

  // Method overriding
  void displayOccupation() const override
  {
    cout << getName() << " is a FAANG engineer at " << company << "." << endl;
  }
};

class StartupEngineer : public SoftwareEngineer
{
public:
  // Constructor using base class constructor
  StartupEngineer(string name, int age, string company) : SoftwareEngineer(name, age, company) {}

  // Method overriding
  void displayOccupation() const override
  {
    cout << getName() << " is a startup engineer at " << company << "." << endl;
  }
};

int main()
{
  // Create objects for each class
  Person person("John Doe", 30);
  Employee employee("Jane Doe", 28, "ABC Corp");
  SoftwareEngineer engineer("Bob Smith", 25, "XYZ Inc");
  FAANGEngineer faangEngineer("Alice Johnson", 27, "Google");
  StartupEngineer startupEngineer("Tom Brown", 26, "TechStartup");

  // Demonstrate Polymorphism and method overriding
  person.displayOccupation();
  employee.displayOccupation();
  engineer.displayOccupation();
  faangEngineer.displayOccupation();
  startupEngineer.displayOccupation();

  // Use friend class to view occupation
  OccupationViewer::viewOccupation(person);
  OccupationViewer::viewOccupation(employee);
  OccupationViewer::viewOccupation(engineer);

  return 0;
}
