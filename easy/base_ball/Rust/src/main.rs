mod base_ball;

pub mod test_base_ball;

fn main() 
{
    // let input: Vec<String> = vec![""]
    let inputs: [&[&str]; 3] = [
        &["5","2","C","D","+"],
        &["5","-2","4","C","D","9","+","+"],
        &["1"],
    ];

    for input in inputs {
        let in_v: Vec<String> = input.iter().map(|&s| s.into()).collect();
        println!("{}", base_ball::calc_points(in_v));
    }
}
