// Success
// Details
// Runtime: 2 ms, faster than 43.15% of Rust online submissions (2,800,170) for Valid Parentheses.
// Memory Usage: 2.2 MB, less than 40.88% of Rust online submissions (2,800,170) for Valid Parentheses.

// 92 / 92 test cases passed.
// 	Status: Accepted
// Runtime: 2 ms
// Memory Usage: 2.2 MB

// Date: 7.11.2022

// Copyright 2022 Bogdan Peta, All Rights Reserved

use std::collections::HashMap;

// In Go if a key does not exists in map map[key] 
// returns nil for that type, 0 for integer in our case
// Rust panics when a key does not exists, so emulate 
// that with unwrap_or


pub fn match_paranthesis(s: String) -> bool
{
    if s.len() < 1 || s.len() > 10000 {
        return false;
    }

    let mut temp: Vec<u8> = Vec::new();

    let pairs: HashMap<u8, u8> = [
        (40,41), (91,93), (123,125),
    ].into();

    for ch in s.bytes() {
        if temp.len() == 0 {
            temp.push(*pairs.get(&ch).unwrap_or(&0));
        } else {    
            let last = temp[temp.len() - 1];
            if ch == last {
                temp.pop();
            } else {
                temp.push(*pairs.get(&ch).unwrap_or(&0));
            }
        }
    }

    temp.len() == 0
}
