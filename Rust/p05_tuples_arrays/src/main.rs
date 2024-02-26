// Tuples - creation & usage
fn create_and_use_tuples() {
    // Creating a tuple
    let person: (String, i32, bool) = ("Alice".to_string(), 30, true);

    // Accessing elements of the tuple
    let name = person.0;
    let age = person.1;
    let is_adult = person.2;

    // Printing tuple elements
    println!("Name: {}", name);
    println!("Age: {}", age);
    println!("Is Adult: {}", is_adult);
}

// Arrays - creation & usage
fn create_and_use_arrays() {
    // Creating an array
    let numbers: [i32; 5] = [1, 2, 3, 4, 5];

    // Accessing elements of the array
    let first_number = numbers[0];
    let last_number = numbers[numbers.len() - 1];

    // Printing array elements
    println!("First Number: {}", first_number);
    println!("Last Number: {}", last_number);

    // Iterating through the array
    for number in numbers.iter() {
        println!("Number: {}", number);
    }
}

fn main() {
    create_and_use_tuples();
    create_and_use_arrays();
}
