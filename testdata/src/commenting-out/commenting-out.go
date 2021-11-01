// want "found dangerous unicode character sequence RIGHT-TO-LEFT-OVERRIDE at position 381" "found dangerous unicode character sequence LEFT-TO-RIGHT-ISOLATE at position 387" "found dangerous unicode character sequence POP-DIRECTIONAL-ISOLATE at position 402"
package main

import "fmt"

func main() {
	isAdmin := false
	isSuperAdmin := false
	isAdmin = isAdmin || isSuperAdmin
	/*‮ } ⁦if (isAdmin)⁩ ⁦ begin admins only */
	fmt.Println("You are an admin.")
	/* end admins only ‮ { ⁦*/
}
