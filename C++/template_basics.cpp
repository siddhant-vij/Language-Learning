#include <iostream>
#include <string>

using namespace std;

// Template function
template <typename T>
T addMe(T a, T b)
{
    return a + b;
}

int main()
{
    // Example usage with integer type
    int intResult = addMe(10, 20);
    cout << "The sumi is: " << intResult << '\n';

    // Example usage with float type
    float floatResult = addMe(10.2f, 20.5f);
    cout << "The sumf is: " << floatResult << '\n';

    // Example usage with string
    string stringResult = addMe(string("Hello, "), string("World!"));
    cout << "The sums is: " << stringResult << '\n';

    return 0;
}
