package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if isEvenNum(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func isEvenNum(n int) bool {
	return (n % 2) == 0
}
