package main

import "fmt"

func main() {
	fmt.Println("LEFT-TO-RIGHT-OVERRIDE: '‭'") // want "found dangerous unicode character sequence LEFT-TO-RIGHT-OVERRIDE"
}
