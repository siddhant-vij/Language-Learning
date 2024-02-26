use std::io;

fn calculate(operand1: f64, operand2: f64, operator: char) -> f64 {
    match operator {
        '+' => operand1 + operand2,
        '-' => operand1 - operand2,
        '*' => operand1 * operand2,
        '/' => operand1 / operand2,
        _ => {
            println!("Invalid operator");
            0.0
        }
    }
}

fn main() {
    println!("Enter first number: ");
    let mut num1 = String::new();
    io::stdin().read_line(&mut num1).expect("Failed to read line");
    let num1: f64 = num1.trim().parse().expect("Please enter a number");

    println!("Enter operator (+, -, *, /): ");
    let mut operator = String::new();
    io::stdin().read_line(&mut operator).expect("Failed to read line");
    let operator: char = operator.trim().chars().next().expect("Please enter an operator");

    println!("Enter second number: ");
    let mut num2 = String::new();
    io::stdin().read_line(&mut num2).expect("Failed to read line");
    let num2: f64 = num2.trim().parse().expect("Please enter a number");

    let result = calculate(num1, num2, operator);
    println!("Result: {}", result);
}
