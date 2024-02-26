// Define an enum to represent different product categories
enum ProductCategory {
    Electronics,
    Clothing,
    Books,
}

// Implement a method on the enum to get the tax rate based on the product category
impl ProductCategory {
    fn tax_rate(&self) -> f64 {
        match self {
            ProductCategory::Electronics => 0.15,
            ProductCategory::Clothing => 0.10,
            ProductCategory::Books => 0.05,
        }
    }
}

enum WebEvent {
    PageLoad,
    PageUnload,
    KeyPress(char),
    Paste(String),
    Click { x: i64, y: i64 },
}

impl WebEvent {
    fn description(&self) -> String {
        if let WebEvent::PageLoad = self {
            return String::from("Page loaded");
        } else if let WebEvent::PageUnload = self {
            return String::from("Page unloaded");
        } else if let WebEvent::KeyPress(c) = self {
            return format!("Key pressed: {}", c);
        } else if let WebEvent::Paste(s) = self {
            return format!("Pasted: {}", s);
        } else if let WebEvent::Click { x, y } = self {
            return format!("Clicked at: ({}, {})", x, y);
        } else {
            return String::from("Unknown event");
        }
    }
}

enum Option<T> {
    Some(T),
    None,
}

impl Option<i32> {
    fn description(&self) -> String {
        match self {
            Option::Some(value) => format!("Got a value: {}", value),
            Option::None => String::from("Got None"),
        }
    }
}

fn main() {
    // Create instances of the ProductCategory enum
    let electronics = ProductCategory::Electronics;
    let clothing = ProductCategory::Clothing;
    let books = ProductCategory::Books;

    // Calculate and print the tax rates for each product category
    println!("Tax rate for Electronics: {:.2}", electronics.tax_rate());
    println!("Tax rate for Clothing: {:.2}", clothing.tax_rate());
    println!("Tax rate for Books: {:.2}", books.tax_rate());

    let load = WebEvent::PageLoad;
    let unload = WebEvent::PageUnload;
    let keypress = WebEvent::KeyPress('x');
    let paste = WebEvent::Paste("Hello, world!".to_string());
    let click = WebEvent::Click { x: 20, y: 80 };

    println!("{}", load.description());
    println!("{}", unload.description());
    println!("{}", keypress.description());
    println!("{}", paste.description());
    println!("{}", click.description());

    let some_value = Option::Some(42);
    let no_value: Option<i32> = Option::None;

    println!("{}", some_value.description());
    println!("{}", no_value.description());
}
