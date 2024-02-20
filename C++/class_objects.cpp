#include <iostream>
#include <string>

using namespace std;

// Declaration of the Person class
class Person
{
private:
  // Private member variables
  string _name;
  int _age;
  // Any constructor can be disabled (no object instantiation allowed with that constructor) if it is placed in the private section. 

public:
  // Default constructor
  Person();

  // Parameterized constructor
  Person(string name, int age);

  // Copy constructor
  Person(const Person &other);

  // Copy assignment operator
  Person &operator=(const Person &other);

  // Destructor
  ~Person();

  // Public getters (const qualified methods)
  string getName() const;
  int getAge() const;

  // Public setters
  void setName(string newName);
  void setAge(int newAge);
};

// Implementation of the Person class

// Default constructor
Person::Person() : _name("Unknown"), _age(0)
{
  cout << "Default constructor called" << endl;
}

// Parameterized constructor
Person::Person(string name, int age) : _name(name), _age(age)
{
  cout << "Parameterized constructor called" << endl;
}

// Copy constructor
Person::Person(const Person &other) : _name(other._name), _age(other._age)
{
  cout << "Copy constructor called" << endl;
}

// Copy assignment operator
Person &Person::operator=(const Person &other)
{
  cout << "Copy assignment operator called" << endl;
  if (this != &other)
  { // Protect against invalid self-assignment
    this->_name = other._name;
    this->_age = other._age;
  }
  return *this;
}

// Destructor
Person::~Person()
{
  cout << "Destructor called for " << this << endl;
}

// Getter for name
string Person::getName() const
{
  return _name;
}

// Getter for age
int Person::getAge() const
{
  return _age;
}

// Setter for name
void Person::setName(string newName)
{
  _name = newName;
}

// Setter for age
void Person::setAge(int newAge)
{
  _age = newAge;
}

// Main function to demonstrate usage of the Person class
int main()
{
  // Using the default constructor
  Person person1;
  cout << "Person 1: " << person1.getName() << ", " << person1.getAge() << " years old" << endl;
  cout << "Address of Person 1: " << &person1 << endl;

  // Using the parameterized constructor
  const Person person2("Alice", 30);
  cout << "Person 2: " << person2.getName() << ", " << person2.getAge() << " years old" << endl;
  cout << "Address of Person 2: " << &person2 << endl;

  // Using the copy constructor
  const Person person3(person2);
  cout << "Person 3 (copy of Person 2): " << person3.getName() << ", " << person3.getAge() << " years old" << endl;
  cout << "Address of Person 3: " << &person3 << endl;

  // Using the copy assignment operator
  const Person person4 = person2;
  cout << "Person 4 (after assignment from Person 2): " << person4.getName() << ", " << person4.getAge() << " years old" << endl;
  cout << "Address of Person 4: " << &person4 << endl;

  // Setting values using setters
  // Since setters not defined with const, setters can't be used with const objects: person2, 3 & 4.
  person1.setName("Bob");
  person1.setAge(25);
  cout << "Person 1 (after setters): " << person1.getName() << ", " << person1.getAge() << " years old" << endl;

  return 0;
}
