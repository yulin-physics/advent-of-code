fn part_a(nums: &Vec<i64>) -> i64 {
    let mut count = 0;
    for i in 1..nums.len() {
        if nums[i] - nums[i - 1] > 0 {
            count += 1;
        }
    }
    count
}

fn part_b(nums: &Vec<i64>) -> i64 {
    let mut count = 0;
    for i in 3..nums.len() {
        if nums[i] - nums[i - 3] > 0 {
            count += 1;
        }
    }
    count
}

fn parse_nums(input: &str) -> Vec<i64> {
    input
        .trim()
        .split("\n")
        .map(|l| l.parse::<i64>().unwrap())
        .collect()
}

fn main() {
    let nums = parse_nums(include_str!("inputs/day_01.txt"));

    println!("Part a: {}\nPart b: {}", part_a(&nums), part_b(&nums));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_a() {
        let nums = parse_nums(include_str!("inputs/test_01.txt"));
        assert_eq!(part_a(&nums), 7);
    }

    #[test]
    fn test_part_b() {
        let nums = parse_nums(include_str!("inputs/test_01.txt"));
        assert_eq!(part_b(&nums), 5);
    }
}
