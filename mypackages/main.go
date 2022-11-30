package arithmetic

import "fmt"

func main() {
	fmt.Println("Main")

	x := 3
	fmt.Printf("%d is prime: %v", x, IsPrime(x))
}
