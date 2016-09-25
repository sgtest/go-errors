package foo

func AAA() int {
	return "hello"
}

func BBB() string {
	return AAA()
}
