// Success
// Details
// Runtime: 31 ms, faster than 14.34% of Rust online submissions for Group Anagrams.
// Memory Usage: 6 MB, less than 15.97% of Rust online submissions for Group Anagrams.

// 118 / 118 test cases passed.
// 	Status: Accepted
// Runtime: 31 ms
// Memory Usage: 6 MB

// Date: 9.11.2022

// Copyright 2022 Bogdan Peta, All Rights Reserved

use std::collections::HashMap;

pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>>
{
    let mut acc: HashMap<String, Vec<String>> = HashMap::new();

    for mut w in strs {

        w = w.to_lowercase();
        let mut str_w: Vec<char> = w.chars().collect::<Vec<char>>();
        str_w.sort_by(|s1, s2| s1.cmp(s2));

        let w1 = str_w.into_iter().collect::<String>();
        
        let mut a: Vec<String> = acc.get(&w1).unwrap_or(&vec![]).to_vec();

        a.push(w);
        acc.insert(w1, a);        
    }

    // sort the map to always pass the tests
    let mut sa: Vec::<String> = acc.keys().cloned().collect();
    // sa.sort_by_key(|a| a.to_lowercase());
    sa.sort();

    let mut rez: Vec<Vec<String>> = vec![];

    for val in sa {
        rez.push(acc.get(&val).unwrap().to_vec());
    }
    
    return rez;
}
