
#[path = "base_ball.rs"] pub mod base_ball;

#[cfg(test)]

// use super::*;

#[test]
fn calc_points_test1() {
    let input: Vec<String> = ["5","2","C","D","+"].iter().map(|&s| s.into()).collect();

    let rez = base_ball::calc_points(input);

    assert_eq!(rez, 30);
}

#[test]
fn calc_points_test2() {
    let input:Vec<String> = ["5","-2","4","C","D","9","+","+"].iter().map(|&s| s.into()).collect();

    let rez = base_ball::calc_points(input);

    assert_eq!(rez, 27);
}

#[test]
fn calc_points_test3() {
    let input: Vec<String> = ["1"].iter().map(|&s| s.into()).collect();

    let rez = base_ball::calc_points(input);

    assert_eq!(rez, 1);
}
