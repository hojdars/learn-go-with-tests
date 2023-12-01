package iteration

func Iterate(letter string, number int) (result string) {
	for i := 0; i < number; i++ {
		result += letter
	}
	return
}
