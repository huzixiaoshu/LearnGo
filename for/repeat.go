package iteration

func Repeat(character string, num int) string {
	repeated := ""
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}
