// Success
// Details
// Runtime: 0 ms, faster than 100.00% of Rust online submissions for Baseball Game.
// Memory Usage: 2.3 MB, less than 12.00% of Rust online submissions for Baseball Game.

// 39 / 39 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.3 MB

// Copyright 2022 Bogdan Peta

pub fn calc_points(operations: Vec<String>) -> i32 
{
    let mut out: Vec<i32> = vec![];

    for ch in operations {
        let val0 = ch.parse::<i32>();

        match val0 {
            Ok(val) => {
                out.push(val);
            },
            Err(_e) => {
                if ch == "D" {
                    let crt = out.pop().unwrap();
                    out.push(crt);
                    out.push(crt * 2);
                } else if ch == "C" {
                    out.pop().unwrap();
                } else if ch == "+"{
                    let x1 = out.pop().unwrap();
                    let x0 = out.pop().unwrap();
                    out.push(x0);
                    out.push(x1);
                    out.push(x0 + x1);
                }
            },
        }
    }

    let mut sum = 0;

    for elem in out {
        sum += elem;
    }

    sum
}
