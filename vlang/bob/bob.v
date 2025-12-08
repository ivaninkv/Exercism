module main

fn has_letters(s string) bool {
	for ch in s.runes() {
		if ch.str().to_upper() != ch.str().to_lower() {
			return true
		}
	}
	return false
}

fn response(hey_bob string) string {
	trimmed := hey_bob.trim_space()

	if trimmed == '' {
		return 'Fine. Be that way!'
	}

	is_question := trimmed.ends_with('?')
	is_shouting := trimmed.to_upper() == trimmed && has_letters(trimmed)

	match true {
		is_shouting && is_question {
			return "Calm down, I know what I'm doing!"
		}
		is_shouting {
			return 'Whoa, chill out!'
		}
		is_question {
			return 'Sure.'
		}
		else {
			return 'Whatever.'
		}
	}
}

fn main() {
	s := '1, 2, 3 GO!'
	println(response(s))
	println(has_letters(s))
}
