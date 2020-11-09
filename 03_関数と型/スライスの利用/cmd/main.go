package main

func main() {
	ns := []int{19, 86, 1, 12}
	var sum int
	for _, n := range ns {
		sum += n
	}
	println(sum)
}
