module main

fn is_isogram(word string) bool {
	word_lower := word.to_lower()
	for c in word_lower {
		if c.is_letter() && word_lower.count(c.ascii_str()) > 1 {
			return false
		}
	}
	return true
}

fn main() {
	println(is_isogram('eleven'))
}
