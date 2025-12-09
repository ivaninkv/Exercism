module main

fn is_pangram(phrase string) bool {
	letter_runes := phrase.to_lower().runes().filter(it >= `a` && it <= `z`)
	mut unique_set := map[rune]bool{}
	for r in letter_runes {
		unique_set[r] = true
	}

	return unique_set.len == 26
}

fn is_pangram_best(phrase string) bool {
	return 'abcdefghijklmnopqrstuvwxyz'.split('')
		.all(phrase.to_lower().contains(it))
}

fn main() {
	println(is_pangram('the quick brown fox jumps over the lazy dog'))
	println(is_pangram_best('the quick brown fox jumps over the lazy dog'))
}
