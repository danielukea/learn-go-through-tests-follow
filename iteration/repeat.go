package iteration

func Repeat(character string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + character
	}
	return repeated
}
