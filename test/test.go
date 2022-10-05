package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')

	// inter
	// args := os.Args
	// s := ""
	// have := false
	// if len(args) == 3 {
	// 	for _, v := range args[1] {
	// 		have = false
	// 		for _, k := range args[2] {
	// 			if v == k {
	// 				for _, f := range s {
	// 					if k == f {
	// 						have = true
	// 						break
	// 					}
	// 				}
	// 				if !have {
	// 					s += string(k)
	// 				}
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(s)
}
