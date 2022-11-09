mod group_anagrams;

mod test_group_anagrams;

fn main() 
{

    let inputs : Vec<Vec<&str>> = vec![
		vec!["eat","tea","tan","ate","naT","bat", "Eat", "TEA"],
		vec![""],
		vec!["a"],
		vec!["", ""],
		vec!["c", "c"],
		vec!["h", "h", "h"],
		vec!["bdddddddddd","bbbbbbbbbbc"],
		vec!["","b",""],
		vec!["tea","","eat","","tea","", "", "Tea"]
    ];

    for input in inputs {
		let input0 = input.iter().map(|s| s.to_string()).collect::<Vec<String>>();

		print!("{:?}: ", input0);

        let out = group_anagrams::group_anagrams(input0);

        println!("{:?}", out);
    }
}
