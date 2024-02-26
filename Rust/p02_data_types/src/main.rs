fn signed_integers() {
    println!("i8 size: {} bytes, min: {}, max: {}", std::mem::size_of::<i8>(), i8::MIN, i8::MAX);
    println!("i16 size: {} bytes, min: {}, max: {}", std::mem::size_of::<i16>(), i16::MIN, i16::MAX);
    println!("i32 size: {} bytes, min: {}, max: {}", std::mem::size_of::<i32>(), i32::MIN, i32::MAX);
    println!("i64 size: {} bytes, min: {}, max: {}", std::mem::size_of::<i64>(), i64::MIN, i64::MAX);
    println!("i128 size: {} bytes, min: {}, max: {}", std::mem::size_of::<i128>(), i128::MIN, i128::MAX);
    println!("isize size: {} bytes, min: {}, max: {}", std::mem::size_of::<isize>(), isize::MIN, isize::MAX);
}

fn unsigned_integers() {
    println!("u8 size: {} bytes, min: {}, max: {}", std::mem::size_of::<u8>(), u8::MIN, u8::MAX);
    println!("u16 size: {} bytes, min: {}, max: {}", std::mem::size_of::<u16>(), u16::MIN, u16::MAX);
    println!("u32 size: {} bytes, min: {}, max: {}", std::mem::size_of::<u32>(), u32::MIN, u32::MAX);
    println!("u64 size: {} bytes, min: {}, max: {}", std::mem::size_of::<u64>(), u64::MIN, u64::MAX);
    println!("u128 size: {} bytes, min: {}, max: {}", std::mem::size_of::<u128>(), u128::MIN, u128::MAX);
    println!("usize size: {} bytes, min: {}, max: {}", std::mem::size_of::<usize>(), usize::MIN, usize::MAX);
}

fn floating_point_numbers() {
    println!("f32 size: {} bytes, min: {}, max: {}", std::mem::size_of::<f32>(), f32::MIN, f32::MAX);
    println!("f64 size: {} bytes, min: {}, max: {}", std::mem::size_of::<f64>(), f64::MIN, f64::MAX);
}

fn boolean_data_types() {
    println!("bool size: {} bytes, values: {}, {}", std::mem::size_of::<bool>(), false, true);
}

fn character_data_types() {
    println!("char size: {} bytes", std::mem::size_of::<char>());
}

fn main() {
    signed_integers();
    unsigned_integers();
    floating_point_numbers();
    boolean_data_types();
    character_data_types();
}

// *** Overflow and Underflow wrt Data Types