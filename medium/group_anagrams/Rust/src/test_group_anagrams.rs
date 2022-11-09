
#[path = "group_anagrams.rs"] pub mod group_anagrams;

#[cfg(test)]

// use super::*;

#[test]
fn group_anagrams_1()
{
    let test = vec![vec!["bat"], vec!["eat", "tea", "ate"], vec!["tan", "nat"], ];

    let input0: Vec<&str> = vec!["eat","tea","tan","ate","nat","bat"];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_2()
{
    let test = vec![vec![""]];

    let input0: Vec<&str> = vec![""];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_3()
{
    let test = vec![vec!["a"]];

    let input0: Vec<&str> = vec!["a"];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_4()
{
    let test = vec![vec!["",""]];

    let input0: Vec<&str> = vec!["", ""];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_5()
{
    let test = vec![vec!["c", "c"]];

    let input0: Vec<&str> = vec!["c", "c"];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_6()
{
    let test = vec![vec!["h", "h", "h"]];

    let input0: Vec<&str> = vec!["h", "h", "h"];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_7()
{
    let test = vec![vec!["bbbbbbbbbbc"], vec!["bdddddddddd"]];

    let input0: Vec<&str> = vec!["bdddddddddd","bbbbbbbbbbc"];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_8()
{
    let test = vec![vec!["",""], vec!["b"]];

    let input0: Vec<&str> = vec!["", "b", ""];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}

#[test]
fn group_anagrams_9()
{
    let test = vec![vec!["", "", ""], vec!["tea", "eat", "tea"]];

    let input0: Vec<&str> = vec!["tea","","eat","","tea",""];

    let input: Vec<String> = input0.iter().map(|s| s.to_string()).collect::<Vec<String>>();
    
    let out = group_anagrams::group_anagrams(input);

    assert_eq!(out, test);
}
