// Success
// Details
// Runtime: 0 ms, faster than 100.00% of Rust online submissions for Permutations.
// Memory Usage: 2.2 MB, less than 59.85% of Rust online submissions for Permutations.

// 26 / 26 test cases passed.
// 	Status: Accepted
// Runtime: 0 ms
// Memory Usage: 2.2 MB

// Date: 7.11.2022

// Copyright 2022 Bogdan Peta, All Rights Reserved


pub fn generate_permutation(a: &mut Vec<i32>, size: usize, out: &mut Vec<Vec<i32>>)
{
    if size == 1 {
        out.push(a.to_vec());
    }

    for i in 0..size {
        generate_permutation(a, size - 1, out);
        if i < size - 1 {
            if size % 2 == 1 {
                let temp = a[0];
                a[0] = a[size - 1];
                a[size-1] = temp;
            } else {
                let temp = a[i];
                a[i] = a[size - 1];
                a[size-1] = temp;
            }
        }
    }
}


pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> 
{
    let mut out: Vec<Vec<i32>> = vec![];

    let mut a: Vec<i32> = nums;

    let sz = a.len();

    generate_permutation(&mut a, sz, &mut out);
    
    return out;

}