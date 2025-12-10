pub fn is_armstrong_number(num: u32) -> bool {
    let num_str = num.to_string();
    let str_len = num_str.len();
    let mut sum = 0;
    for c in num_str.chars() {
        sum += c.to_digit(10).unwrap().pow(str_len as u32);
    }
    sum == num
}
