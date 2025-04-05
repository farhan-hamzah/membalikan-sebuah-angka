package main
import "fmt"

func angkaTerbalik(n, hasil int) int {
	if n == 0 {
		return hasil
	}
	hasil = hasil*10 + n%10
	return angkaTerbalik(n/10, hasil) 
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print(angkaTerbalik(num, 0)) 
 }