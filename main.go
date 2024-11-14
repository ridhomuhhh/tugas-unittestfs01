package main

import (
	"fmt"
	ewallet "unit-test/app"
)

func main() {
	fmt.Println(ewallet.Run([]string{"deposit", "withdraw", "deposit", "withdraw"}))
	fmt.Println(ewallet.Run([]string{"deposit", "withdraw", "deposit"}))
	fmt.Println(ewallet.Run([]string{"deposit"}))
	fmt.Println(ewallet.Run([]string{"deposit"}))

}
