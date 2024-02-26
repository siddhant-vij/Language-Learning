// Function demonstrating if else
fn if_else_demo(number: i32) {
    if number < 0 {
        println!("Number is negative");
    } else if number == 0 {
        println!("Number is zero");
    } else {
        println!("Number is positive");
    }
}

// Function demonstrating match
fn match_demo(day: &str) {
    match day {
        "Monday" => println!("It's Monday"),
        "Tuesday" | "Wednesday" | "Thursday" | "Friday" => println!("It's a weekday"),
        "Saturday" | "Sunday" => println!("It's the weekend"),
        _ => println!("Invalid day"),
    }
}

// Function demonstrating loops
fn loop_demo() {
    let mut count = 0;

    loop {
        println!("Current count: {}", count);
        count += 1;
        if count == 5 {
            break;
        }
    }

    for i in 1..=5 {
        println!("Value: {}", i);
    }

    let mut while_count = 0;
    while while_count < 5 {
        println!("While count: {}", while_count);
        while_count += 1;
    }
}

// Call the functions to demonstrate the concepts
fn main() {
    if_else_demo(-5);
    if_else_demo(0);
    if_else_demo(5);

    match_demo("Monday");
    match_demo("Friday");
    match_demo("Sunday");

    loop_demo();
}
