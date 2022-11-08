
mod permutations;

mod test_permutations;

fn main() 
{

    let inputs : Vec<Vec<i32>> = vec![
		vec![1,2,3],
		vec![0,1],
		vec![1],
		vec![1,2,3,4],
		vec![1,2,3,4,5],
		vec![1,2,3,4,5,6],
		vec![1,2,3,4,5,6,7],
		vec![1,2,3,4,5,6,7,8],
    ];

    for input in inputs {
        let out = permutations::permute(input);

        println!("{:?}: {}", out, out.len());
    }
}
