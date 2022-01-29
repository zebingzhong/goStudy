package main

func main() {
	mVal := foo(666)
	println(*mVal, mVal)
}

func foo(argVal int) *int {
	var fooVal1 int = 11
	var fooVal2 int = 22
	var fooVal3 int = 33

	for i := 0; i < 1; i++ {
		println(&fooVal1, &fooVal2, &fooVal3)
	}
	return &fooVal2
}
