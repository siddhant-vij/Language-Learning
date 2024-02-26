fn main() {
    // Immutable variable
    let num1 = 5;
    // num1 = 10; // Not allowed to reassign
    println!("Immutable variable: {}", num1);

    // Mutable variable
    let mut num2 = 10;
    num2 += 5;
    println!("Mutable variable: {}", num2);

    // Constant variable
    const CONSTANT_NUM: i32 = 100;
    println!("Constant value: {}", CONSTANT_NUM);
}

// Constants are always fixed in size, and known at compile time. Constants may not be of a type that requires allocation on the heap, since they’re not known at compile time.

// Constants can only be set to a constant expression, not to the result of a function call or anything that could only be determined at runtime.