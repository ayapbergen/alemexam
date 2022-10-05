package main

import (
	"fmt"
	"piscine"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.SortWordArr(result)

	fmt.Println(result)

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
