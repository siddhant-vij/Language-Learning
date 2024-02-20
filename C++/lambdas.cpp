#include <iostream>
#include <vector>
#include <algorithm>
#include <functional>

using namespace std;

int main() {
    // A simple lambda expression that adds two numbers
    auto add = [](int a, int b) -> int { return a + b; };
    cout << "The sum of 5 and 3 is " << add(5, 3) << endl;

    // A lambda that captures local variables by value
    int factor = 2;
    auto multiplyByFactor = [factor](int n) { return n * factor; };
    cout << "4 multiplied by factor (2) is " << multiplyByFactor(4) << endl;

    // Modifying captured local variables using mutable keyword
    auto counter = [count = 0]() mutable { return ++count; };
    cout << "Counter: " << counter() << endl;
    cout << "Counter: " << counter() << endl;

    // Using lambdas with the standard algorithm library
    vector<int> numbers = {1, 2, 3, 4, 5};
    cout << "Original vector: ";
    for (int n : numbers) {
        cout << n << " ";
    }
    cout << endl;

    // Multiply each number in the vector by the factor using a lambda expression
    for_each(numbers.begin(), numbers.end(), [factor](int &n) { n *= factor; });
    cout << "After multiplying by factor: ";
    for (int n : numbers) {
        cout << n << " ";
    }
    cout << endl;

    // Using a lambda to sort a vector in descending order
    sort(numbers.begin(), numbers.end(), [](int a, int b) { return a > b; });
    cout << "Vector sorted in descending order: ";
    for (int n : numbers) {
        cout << n << " ";
    }
    cout << endl;

    // Storing a lambda in a function
    function<int(int, int)> subtract = [](int a, int b) { return a - b; };
    cout << "The result of subtracting 5 from 10 is " << subtract(10, 5) << endl;

    return 0;
}
