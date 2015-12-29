package utils

func Reverse(arg string) string {
	buf := []rune(arg)
	for i, j := 0, len(buf) - 1; i < len(buf) / 2; i, j = i + 1, j - 1 {
		buf[i], buf[j] = buf[j], buf[i]
	}

	return string(buf)
}
