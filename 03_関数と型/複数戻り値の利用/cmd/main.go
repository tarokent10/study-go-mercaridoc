package main

func main() {
	n, m := swap(10, 20)
	println(n, m)
}
func swap(n, m int) (int, int) {
	return m, n
}
