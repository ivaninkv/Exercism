package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	result := make([]string, len(rhyme))
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			result[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
		result[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}
	return result
}
