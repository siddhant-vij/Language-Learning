use rand::Rng;
use std::{collections::HashSet, io};

fn main() {
    let mut guessing_sample_space = String::new();

    println!("How many numbers should be present in the guessing sample space?: ");
    io::stdin().read_line(&mut guessing_sample_space).unwrap();
    let guessing_sample_space: u32 = guessing_sample_space
        .trim()
        .parse()
        .expect("Please type a positive number!");

    if guessing_sample_space == 0 {
        panic!("Please type a number > 0");
    }

    let mut rng = rand::thread_rng();
    let mut numbers_set: HashSet<i32> = HashSet::new();

    while numbers_set.len() < guessing_sample_space as usize {
        let random_num = rng.gen_range(1..=100);
        numbers_set.insert(random_num);
    }

    let mut correct_guesses = 0;
    let mut num_guesses = 0;
    println!("Here are the numbers: {:?}", numbers_set);

    loop {
        println!("Enter your guess:");
        let mut guess = String::new();

        io::stdin().read_line(&mut guess).unwrap();
        let guess: i32 = guess.trim().parse().expect("Please type a number!");

        num_guesses += 1;

        if numbers_set.contains(&guess) {
            correct_guesses += 1;
            println!("Correct guess!");

            if correct_guesses == guessing_sample_space {
                break;
            } else {
                continue;
            }
        } else {
            println!("Incorrect guess, try again.");
        }
    }

    println!(
        "You guessed all numbers correctly! Total number of guesses: {}",
        num_guesses
    );
}
