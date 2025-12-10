module main

const scores = {
	'aeioulnrst': 1
	'dg':         2
	'bcmp':       3
	'fhvwy':      4
	'k':          5
	'jx':         8
	'qz':         10
}

fn score(word string) int {
	mut res := 0

	for c in word.to_lower() {
		for k, v in scores {
			if k.contains(c.ascii_str()) {
				res += v
			}
		}
	}

	return res
}

fn main() {
	println(score('cabbage'))
}
