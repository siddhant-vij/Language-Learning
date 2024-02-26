// An integer division that doesn't `panic!`
fn checked_division_option(dividend: i32, divisor: i32) -> Option<i32> {
    if divisor == 0 {
        // Failure is represented as the `None` variant
        None
    } else {
        // Result is wrapped in a `Some` variant
        Some(dividend / divisor)
    }
}

// This function handles a division that may not succeed
fn try_division_option(dividend: i32, divisor: i32) {
    // `Option` values can be pattern matched, just like other enums
    match checked_division_option(dividend, divisor) {
        None => println!("{} / {} failed!", dividend, divisor),
        Some(quotient) => {
            println!("{} / {} = {}", dividend, divisor, quotient)
        },
    }
}

fn option_demo() {
    try_division_option(4, 2);
    try_division_option(1, 0);
}

// An integer division that doesn't `panic!`
fn checked_division_result(dividend: i32, divisor: i32) -> Result<i32, &'static str> {
    if divisor == 0 {
        // Failure is represented as the `Err` variant
        Err("Division by zero is not allowed")
    } else {
        // Result is wrapped in an `Ok` variant
        Ok(dividend / divisor)
    }
}

// This function handles a division that may not succeed
fn try_division_result(dividend: i32, divisor: i32) {
    // `Result` values can be pattern matched, just like other enums
    match checked_division_result(dividend, divisor) {
        Err(e) => println!("{} / {} failed! Error: {}", dividend, divisor, e),
        Ok(quotient) => {
            println!("{} / {} = {}", dividend, divisor, quotient)
        },
    }
}

fn result_demo() {
    try_division_result(4, 2);
    try_division_result(1, 0);
}

fn main() {
    option_demo();
    result_demo();
}
