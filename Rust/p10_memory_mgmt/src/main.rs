// This function takes ownership of the heap allocated memory
fn own_box(c: Box<i32>) {
    println!("Destroying a box that contains {}", c);

    // `c` is destroyed and the memory freed
}

fn borrow_box(b: &Box<i32>) {
    println!("Borrowing a box that contains {}", b);
}

fn ownership_demo() {
    // _Stack_ allocated integer
    let x = 5u32;

    // Copy `x` into `y` - no resources are moved
    let y = x;

    // Both values can be independently used
    println!("x is {}, and y is {}", x, y);

    // `a` is a pointer to a _heap_ allocated integer
    let a = Box::new(5i32);

    println!("a contains: {}", a);

    // Move `a` into `b`
    let b = a;
    // The pointer address of `a` is copied (not the data) into `b`.
    // Both are now pointers to the same heap allocated data, but `b` now owns it.
    
    // Error! `a` can no longer access the data, because it no longer owns the heap memory
    // println!("a contains: {}", a);
    // TODO ^ Try uncommenting this line

    // This function own_box takes ownership of the heap allocated memory from `b` while borrow_box borrows `b`
    borrow_box(&b);
    own_box(b); // b available here after borrow_box call because it was not moved there - just borrowed

    // Since the heap memory has been freed at this point, this action would result in dereferencing freed memory, but it's forbidden by the compiler
    // Error! Same reason as the previous Error
    // println!("b contains: {}", b);
    // TODO ^ Try uncommenting this line
}


// This function takes ownership of a box and destroys it
fn eat_box_i32(boxed_i32: Box<i32>) {
    println!("Destroying box that contains {}", boxed_i32);
}

// This function borrows an i32
fn borrow_i32(borrowed_i32: &i32) {
    println!("This int is: {}", borrowed_i32);
    // borrowed_i32 += 5; // Not allowed as this is immutable borrow
}

fn mut_borrow_i32(borrowed_i32: &mut Box<i32>) {
    println!("This int is: {}", borrowed_i32);
    *borrowed_i32 = Box::new(4);
}

fn borrowing_demo() {
    // Create a boxed i32 in the heap, and a i32 on the stack
    // Remember: numbers can have arbitrary underscores added for readability
    // 5_i32 is the same as 5i32
    let mut boxed_i32 = Box::new(5_i32);
    let stacked_i32 = 6_i32;

    // Borrow the contents of the box. Ownership is not taken, so the contents can be borrowed again.
    mut_borrow_i32(&mut boxed_i32);
    borrow_i32(&boxed_i32);
    borrow_i32(&stacked_i32);

    {
        // Take a reference to the data contained inside the box
        let _ref_to_i32: &i32 = &boxed_i32;

        // Error!
        // Can't destroy `boxed_i32` while the inner value is borrowed later in scope.
        // eat_box_i32(boxed_i32);
        // TODO ^ Try uncommenting this line

        // Attempt to borrow `_ref_to_i32` after inner value is destroyed (if above line uncommented) & this is why we can't destroy `boxed_i32` in above line
        borrow_i32(_ref_to_i32);
        // `_ref_to_i32` goes out of scope and is no longer borrowed.
    }

    // `boxed_i32` can now give up ownership to `eat_box` and be destroyed
    eat_box_i32(boxed_i32);
}

fn main() {
    ownership_demo();
    borrowing_demo();
}
