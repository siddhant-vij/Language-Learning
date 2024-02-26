struct Grocery {
    name: String,
    quantity: u32,
}

fn print_quantity(grocery: &Grocery) {
    println!("{}", grocery.quantity);
}

fn print_name(grocery: &Grocery) {
    println!("{}", grocery.name);
}

fn main() {
    let grocery = Grocery {
        name: String::from("milk"),
        quantity: 2,
    };
    print_quantity(&grocery);
    print_name(&grocery);
}
