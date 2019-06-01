package reverse

func Reverse(str string) string {
	if len(str) == 1 {
		return str
	}
	return Reverse(str[1:]) + string(str[0])
}
