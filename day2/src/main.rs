use std::fs::File;
use std::io::{BufRead, BufReader};
fn main() {
    let file_path = "src/input.txt";
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    let mut valid_count = 0;

    for (line_num, line) in reader.lines().enumerate() {
        if let Ok(line) = line {
            println!("Line {}: {}", line_num + 1, line);

            // Split the line into integers
            let line_arr: Vec<i32> = line
                .split_whitespace()
                .filter_map(|s| s.parse::<i32>().ok())
                .collect();

            if traverse(line_arr) {
                valid_count += 1;
                println!("Line {}: Valid", line_num + 1);
            } else {
                println!("Line {}: Invalid", line_num + 1);
            }
        }

        println!("{}", valid_count);
    }
}
fn traverse(line_arr: Vec<i32>) -> bool {
    let mut flag = true;
    // Check if the report is safe without removing any levels
    if is_safe(line_arr.clone()) {
        return true;
    }

    // Try removing each level one by one
    for i in 0..line_arr.len() {
        if !flag {
            break;
        }
        let mut new_arr = line_arr.clone();
        new_arr.remove(i);
        if is_safe(new_arr) {
            return true;
        }
        flag = true;
    }

    // If none of the above works, return false
    false
}
fn is_safe(line_arr: Vec<i32>) -> bool {
    let mut is_increasing: Option<bool> = None; // Tracks if the sequence is increasing or decreasing
    let mut valid = true;

    for i in 0..line_arr.len() - 1 {
        // Calculate the difference between consecutive elements
        let difference = line_arr[i + 1] - line_arr[i];

        // Check if the sequence is strictly increasing or decreasing
        if difference > 0 {
            if is_increasing == Some(false) {
                valid = false; // Switch from decreasing to increasing is invalid
                break;
            }
            is_increasing = Some(true);
        } else if difference < 0 {
            if is_increasing == Some(true) {
                valid = false; // Switch from increasing to decreasing is invalid
                break;
            }
            is_increasing = Some(false);
        }

        // Ensure the absolute difference is within the valid range [1, 3]
        if difference.abs() < 1 || difference.abs() > 3 {
            valid = false; // Difference out of range
            break;
        }
    }
    valid
}
