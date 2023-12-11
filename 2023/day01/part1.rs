use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::time::Instant;

fn main() -> io::Result<()> {
    let start_timer = Instant::now();
    let path = Path::new("input.txt");
    let file = File::open(&path)?;
    let reader = io::BufReader::new(file);

    let mut sum = 0;
    for line in reader.lines() {
        let line = line?;
        let mut left: Option<char> = None;
        let mut right: Option<char> = None;
        for ch in line.chars() {
            if ch.is_digit(10) {
                if left.is_none() {
                    left = Some(ch);
                    right = Some(ch);
                } else {
                    right = Some(ch);
                }
            }
        }

        if let (Some(l), Some(r)) = (left, right) {
            sum += l.to_digit(10).unwrap() * 10 + r.to_digit(10).unwrap();
        }
    }

    println!("{}", sum);
    println!("{} ms", start_timer.elapsed().as_micros());

    Ok(())
}