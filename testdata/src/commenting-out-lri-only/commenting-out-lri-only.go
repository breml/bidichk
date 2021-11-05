package main

import "fmt"

func main() {
	isAdmin := false
	isSuperAdmin := false
	isAdmin = isAdmin || isSuperAdmin
	/*‮ } ⁦if (isAdmin)⁩ ⁦ begin admins only */ // want "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE" "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE"
	fmt.Println("You are an admin.")
	/* end admins only ‮ { ⁦*/ // want "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE"
}
