package main

import "fmt"

func main() {
	accessLevel := "user"
	if accessLevel != "user‮ ⁦// Check if admin⁩ ⁦" { // want "found dangerous unicode character sequence POP-DIRECTIONAL-ISOLATE" "found dangerous unicode character sequence RIGHT-TO-LEFT-OVERRIDE" "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE" "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE"
		fmt.Println("You are an admin.")
	}
}
