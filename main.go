package main

func main() {
	a := [...]int{81, 54, 43, 66}
	for i, v := range a {
		println(i, v)
	}
}
