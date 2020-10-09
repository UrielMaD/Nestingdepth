package main

import (
	"fmt"
)

func main() {

	var iteraciones int
	fmt.Scanln(&iteraciones)

	for l := 0; l < iteraciones; l++ {

		var nums string
		fmt.Scanln(&nums)
		depth := int(nums[0]) - 48
		cad := ""

		for i, n := range nums {
			if i == 0 {
				for j := 0; j < int(n-48); j++ {
					cad += "("
				}
				cad += string(n)
			} else {
				if int(n) == int(nums[i-1]) {
					cad += string(n)
				}
				if int(n) > int(nums[i-1]) {
					for j := 0; j < (int(n) - int(nums[i-1])); j++ {
						cad += "("
						depth++
					}
					cad += string(n)
				} else if int(n) < int(nums[i-1]) {
					for j := 0; j < (int(nums[i-1]) - int(n)); j++ {
						cad += ")"
						depth--
					}
					cad += string(n)

				}

			}

		}
		for i := 0; i < depth; i++ {
			cad += ")"
		}

		fmt.Printf("Case #%d: %s\n", l+1, cad)
	}
}
