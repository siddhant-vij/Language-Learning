// Define a trait named Printable
trait Printable {
    // Method signature for printing
    fn print(&self) {
        println!("Printing generic type");
    }
}

// Implement the trait Printable for the type i32
impl Printable for i32 {
    fn print(&self) {
        println!("Printing i32: {}", self);
    }
}

// Define a struct named Person
struct Person {
    name: String,
    age: u32,
}

// Implement the trait Printable for the struct Person
impl Printable for Person {
    fn print(&self) {
        println!("Printing Person - Name: {}, Age: {}", self.name, self.age);
    }
}

impl Printable for f64 {}

fn main() {
    let num = 42;
    num.print(); // Output: Printing i32: 42

    let person = Person {
        name: String::from("Alice"),
        age: 30,
    };
    person.print(); // Output: Printing Person - Name: Alice, Age: 30

    let fnum = 3.14;
    fnum.print(); // Output: Printing generic type
}
