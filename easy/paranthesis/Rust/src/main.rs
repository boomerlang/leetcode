
mod paranthesis;

mod test_paranthesis;

fn main() 
{

    let inputs = [
        "(]",
        "([)]",
         "()",
        "{[]}",
        "()[]{}",
        "()[({})]{[({})]}",
        "([{}])",
        "",
    ];
        
    for input in inputs {
        println!("{}", paranthesis::match_paranthesis(input.to_string()));
    }
}
