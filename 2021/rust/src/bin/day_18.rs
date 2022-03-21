// mutable circular pointers difficult to do in rust
// use arrays to store nodes instead, and use indexes as pointers to nodes

use std::io::{self, Read};

#[derive(Debug)]
enum Tree {
    Pair(Box<Tree>, Box<Tree>),
    Leaf(u32),
}

fn get_input() -> String {
    let mut input = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

fn add_array_tree(
    a_vals: &mut Vec<u32>,
    a_depths: &mut Vec<u32>,
    b_vals: &mut Vec<u32>,
    b_depths: &mut Vec<u32>,
) {
    a_vals.append(b_vals);
    a_depths.append(b_depths);
    for i in 0..a_depths.len() {
        a_depths[i] += 1;
    }
}

fn reduce(vals: &mut Vec<u32>, depths: &mut Vec<u32>) {
    loop {
        let did_explode = explode(vals, depths);
        if did_explode {
            continue;
        };

        let did_split = split(vals, depths);

        if !did_explode && !did_split {
            break;
        }
    }
}

fn split(vals: &mut Vec<u32>, depths: &mut Vec<u32>) -> bool {
    for i in 0..vals.len() {
        let val = vals[i];
        if val < 10 {
            continue;
        }

        let (a, b) = match val % 2 {
            0 => (val / 2, val / 2),
            _ => (val / 2, val / 2 + 1),
        };

        vals[i] = a;
        depths[i] += 1;
        vals.insert(i + 1, b);
        depths.insert(i + 1, depths[i]);

        return true;
    }
    false
}

fn explode(vals: &mut Vec<u32>, depths: &mut Vec<u32>) -> bool {
    for i in 0..depths.len() {
        let depth = depths[i];
        if depth != 4 {
            continue;
        }

        // add left value to left neighbour
        if i != 0 {
            vals[i - 1] += vals[i];
        }

        // add right value to right neighbour
        if i + 2 <= vals.len() - 1 {
            vals[i + 2] += vals[i + 1];
        }

        vals[i] = 0;
        depths[i] = 3;
        vals.remove(i + 1);
        depths.remove(i + 1);

        return true;
    }
    false
}

fn make_array_tree(t: &Tree, depth: u32, vals: &mut Vec<u32>, depths: &mut Vec<u32>) {
    match t {
        Tree::Pair(left, right) => {
            make_array_tree(&left, depth + 1, vals, depths);
            make_array_tree(&right, depth + 1, vals, depths)
        }
        Tree::Leaf(val) => {
            vals.push(*val);
            // starts on depth 0 at a node, so depth for a leaf is $depth -1
            depths.push(depth - 1);
        }
    }
}

fn consume_character(s: &str, c: char) -> &str {
    assert!(s.starts_with(c));
    &s[1..]
}

fn consume_number_literal(s: &str) -> (u32, &str) {
    let c = s.chars().next().unwrap();
    let val = c.to_digit(10).unwrap();
    (val, &s[1..])
}

fn parse(s: &str) -> (Tree, &str) {
    if s.starts_with("[") {
        let s = consume_character(s, '[');
        let (left_subtree, s) = parse(s);
        let s = consume_character(s, ',');
        let (right_subtree, s) = parse(s);
        let s = consume_character(s, ']');
        return (
            Tree::Pair(Box::new(left_subtree), Box::new(right_subtree)),
            s,
        );
    } else {
        let (num, s) = consume_number_literal(s);
        return (Tree::Leaf(num), s);
    }
}

fn score(vals: &mut Vec<u32>, depths: &mut Vec<u32>) -> u32 {
    while vals.len() > 1 {
        for i in 0..depths.len() - 1 {
            if depths[i] == depths[i + 1] {
                vals[i] = 3 * vals[i] + 2 * vals[i + 1];
                vals.remove(i + 1);
                depths.remove(i + 1);
                // depths can be 0
                if depths[i] > 0 {
                    depths[i] -= 1;
                }
                break;
            }
        }
    }

    vals[0]
}

fn main() {
    // let input = get_input();
    // let mut lines = input.lines();

    // let tree = parse(lines.next().unwrap()).0;
    // let mut vals = Vec::new();
    // let mut depths = Vec::new();
    // make_array_tree(&tree, 0, &mut vals, &mut depths);
    // for line in lines {
    //     let other_tree = parse(line).0;
    //     let mut other_vals = Vec::new();
    //     let mut other_depths = Vec::new();

    //     make_array_tree(&other_tree, 0, &mut other_vals, &mut other_depths);
    //     add_array_tree(&mut vals, &mut depths, &mut other_vals, &mut other_depths);
    //     reduce(&mut vals, &mut depths);
    // }
    // let final_score = score(&mut vals, &mut depths);
    // println!("{:?}\n{:?}\n{}", vals, depths, final_score);

    let trees: Vec<Tree> = get_input().lines().map(|line| parse(line).0).collect();
    let mut best_score: u32 = 0;
    for i in 0..trees.len() {
        for j in 0..trees.len() {
            let mut iv = Vec::new();
            let mut id = Vec::new();
            make_array_tree(&trees[i], 0, &mut iv, &mut id);

            let mut jv = Vec::new();
            let mut jd = Vec::new();
            make_array_tree(&trees[j], 0, &mut jv, &mut jd);

            add_array_tree(&mut iv, &mut id, &mut jv, &mut jd);
            reduce(&mut iv, &mut id);
            best_score = best_score.max(score(&mut iv, &mut id));
        }
    }

    println!("{}", best_score);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn can_explode() {
        let (t, _) = parse("[[6,[5,[4,[3,2]]]],1]");
        let mut vals = Vec::new();
        let mut depths = Vec::new();
        make_array_tree(&t, 0, &mut vals, &mut depths);
        explode(&mut vals, &mut depths);
        assert_eq!(vals, vec![6, 5, 7, 0, 3])
    }

    #[test]
    fn can_split() {
        let (t, _) = parse("[1, [10, 1]]");
        let mut vals = Vec::new();
        let mut depths = Vec::new();
        make_array_tree(&t, 0, &mut vals, &mut depths);
        split(&mut vals, &mut depths);
        // assert_eq!(vals, vec![6, 5, 7, 0, 3])
    }
}
