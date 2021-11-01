// want "found dangerous unicode character sequence POP-DIRECTIONAL-ISOLATE at position 373" "found dangerous unicode character sequence RIGHT-TO-LEFT-OVERRIDE at position 349" "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE at position 353"
package main

import "fmt"

func main() {
	accessLevel := "user"
	if accessLevel != "user‮ ⁦// Check if admin⁩ ⁦" {
		fmt.Println("You are an admin.")
	}
}
