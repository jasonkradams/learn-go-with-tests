package iteration

// Repeat takes a character and integer and returns the character n times.
func Repeat(character string, iterations int) string {
	var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}
