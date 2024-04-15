package cell

type Style struct {
	Color string
}

func multiplyRune(r rune, n int) string {
	var str string
	for i := 0; i < n; i++ {
		str += string(r)
	}
	return str
}
