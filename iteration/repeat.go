package iteration

func Repeat(c string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += c
	}
	return result
}
