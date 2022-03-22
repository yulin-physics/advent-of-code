// mutable circular pointers difficult to do in rust
// use arrays to store nodes instead, and use indexes as pointers to nodes

use std::io::{self, Read};

#[derive(Debug, Clone)]
struct VecTree {
    vals: Vec<u32>,
    depths: Vec<u32>,
}

impl VecTree {
    fn parse(s: &str) -> VecTree {
        let mut t = VecTree {
            vals: Vec::new(),
            depths: Vec::new(),
        };

        let mut depth = 0;
        for c in s.chars() {
            match c {
                '[' => depth += 1,
                ',' => (),
                ']' => depth -= 1,
                digit => {
                    t.vals.push(digit.to_digit(10).unwrap());
                    t.depths.push(depth - 1);
                }
            }
        }
        t
    }

    fn reduce(&mut self) {
        loop {
            if !self.explode() && !self.split() {
                break;
            }
        }
    }

    fn split(&mut self) -> bool {
        for i in 0..self.vals.len() {
            let val = self.vals[i];
            if val < 10 {
                continue;
            }

            let (a, b) = match val % 2 {
                0 => (val / 2, val / 2),
                _ => (val / 2, val / 2 + 1),
            };

            self.vals[i] = a;
            self.depths[i] += 1;
            self.vals.insert(i + 1, b);
            self.depths.insert(i + 1, self.depths[i]);

            return true;
        }
        false
    }

    fn explode(&mut self) -> bool {
        for i in 0..self.depths.len() {
            let depth = self.depths[i];
            if depth != 4 {
                continue;
            }

            // add left value to left neighbour
            if i != 0 {
                self.vals[i - 1] += self.vals[i];
            }

            // add right value to right neighbour
            if i + 2 < self.vals.len() {
                self.vals[i + 2] += self.vals[i + 1];
            }

            self.vals[i] = 0;
            self.depths[i] = 3;
            self.vals.remove(i + 1);
            self.depths.remove(i + 1);

            return true;
        }
        false
    }

    fn add(&mut self, other: &VecTree) {
        self.vals.extend(other.vals.iter());
        self.depths.extend(other.depths.iter());
        for i in 0..self.depths.len() {
            self.depths[i] += 1;
        }
    }

    fn score(&self) -> u32 {
        let mut vals = self.vals.clone();
        let mut depths = self.depths.clone();

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
}

fn get_input() -> String {
    let mut input = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

fn main() {
    let input = get_input();
    let mut lines = input.lines();

    let mut tree = VecTree::parse(lines.next().unwrap());
    for line in lines {
        tree.add(&VecTree::parse(line));
        tree.reduce();
    }

    println!("{}", tree.score());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn can_explode() {
        let mut t = VecTree::parse("[[6,[5,[4,[3,2]]]],1]");
        t.explode();
        assert_eq!(t.vals, vec![6, 5, 7, 0, 3])
    }

    #[test]
    fn can_split() {
        let mut t = VecTree::parse("[[[[0,7],4],[7,[[8,4],9]]],[1,1]]");
        t.explode();
        t.split();
        assert_eq!(t.vals, vec![0, 7, 4, 7, 8, 0, 13, 1, 1])
    }
}
