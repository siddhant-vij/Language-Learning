use std::fs::File;
use std::io::{self, Read};

// Define a custom error type
#[derive(Debug)]
enum CustomError {
    FileNotFound,
    IOError(io::Error),
}

// Implement the Display trait for custom error type
impl std::fmt::Display for CustomError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            CustomError::FileNotFound => write!(f, "File not found"),
            CustomError::IOError(ref err) => write!(f, "IO Error: {}", err),
        }
    }
}

// Implement the Error trait for custom error type
impl std::error::Error for CustomError {}

// Function that may result in custom errors
fn read_file_contents(file_path: &str) -> Result<String, CustomError> {
    let mut file = File::open(file_path).map_err(|_| CustomError::FileNotFound)?;

    let mut contents = String::new();
    file.read_to_string(&mut contents).map_err(|err| CustomError::IOError(err))?;

    Ok(contents)
}

// Main function to demonstrate error handling
fn main() {
    let file_path = "example.txt";
    match read_file_contents(file_path) {
        Ok(contents) => println!("File contents: \n{}", contents),
        Err(e) => eprintln!("Error: {}", e),
    }
}
