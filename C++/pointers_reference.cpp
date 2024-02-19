#include <iostream>

using namespace std;

// Swap two integers by value (will not affect the original variables)
void swapByValue(int a, int b)
{
    int temp = a;
    a = b;
    b = temp;
}

// Swap two integers by passing pointers
void swapByPointer(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

// Swap two integers by passing reference
void swapByReference(int &a, int &b)
{
    int temp = a;
    a = b;
    b = temp;
}

int main()
{
    int x = 10, y = 20;

    cout << "Before swapByValue: x = " << x << ", y = " << y << endl;
    swapByValue(x, y); // Call by Value
    cout << "After swapByValue: x = " << x << ", y = " << y << endl;

    cout << "Before swapByPointer: x = " << x << ", y = " << y << endl;
    swapByPointer(&x, &y); // Call by Pointer Reference
    cout << "After swapByPointer: x = " << x << ", y = " << y << endl;

    cout << "Before swapByReference: x = " << x << ", y = " << y << endl;
    swapByReference(x, y); // Call be Reference Variables
    cout << "After swapByReference: x = " << x << ", y = " << y << endl;

    cout << endl
         << "---------------------" << endl;

    int value = 5;     // Normal integer variable
    int *ptr = &value; // Pointer variable that holds the address of 'value'
    int &ref = value;  // Reference to the variable 'value'

    // Print the original value, the pointer, and the reference
    cout << "Original value: " << value << endl;
    cout << "Pointer: " << ptr << endl;
    cout << "Reference: " << ref << endl;

    // Dereference the pointer to get the value it points to
    cout << "\nDereferenced pointer: " << *ptr << endl;

    // Change the value of 'value' using the pointer
    *ptr = 10;
    cout << "Value after dereferencing pointer and changing value: " << value << endl;
    cout << "Pointer: " << ptr << endl;
    cout << "Reference: " << ref << endl;

    // Change the value of 'value' using the reference
    ref = 15;
    cout << "\nValue after changing through reference: " << value << endl;
    cout << "Pointer: " << ptr << endl;
    cout << "Reference: " << ref << endl;

    return 0;
}
