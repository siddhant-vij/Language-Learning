// Define a classic struct
struct ClassicStruct {
    name: String,
    age: u32,
}

// Define a tuple struct
struct TupleStruct(i32, i32);

// Define a unit struct
struct UnitStruct;

// Implement methods for ClassicStruct
impl ClassicStruct {
    // Method to create a new instance
    fn new(name: String, age: u32) -> Self {
        ClassicStruct { name, age }
    }

    // Method to print information about the struct
    fn display_info(&self) {
        println!("Name: {}, Age: {}", self.name, self.age);
    }
}

fn main() {
    // Create instances of the structs
    let classic = ClassicStruct::new(String::from("Alice"), 30);
    let tuple = TupleStruct(10, 20);
    let _unit = UnitStruct;

    // Access and use the classic struct
    classic.display_info();

    // Access and use the tuple struct
    println!("Tuple: ({}, {})", tuple.0, tuple.1);

    // Destructure the tuple struct
    let TupleStruct(x, y) = tuple;
    println!("Tuple: ({}, {})", x, y);

    // Access and use the unit struct
    println!("Unit struct created");
}
