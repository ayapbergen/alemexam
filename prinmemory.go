package piscine

import "fmt"

func PrintMemory(arr [10]int) {
	memory, block, txt := "", 0, ""
	for _, n := range arr {
		h := Hex(n)
		if len(h) < 4 {
			for i := 0; i <= 5-len(h); i++ {
				h += "0"
			}
			h += " 0000"
			if block == 0 {
				memory += h
			} else {
				memory += " " + h
			}

			block++
			if block == 4 {
				memory += "\n"
				block = 0
			}

			if n < 32 {
				txt += "."
			} else {
				txt += string(rune(n))
			}
		}
	}

	fmt.Println(memory)
	fmt.Println(txt)
}

func Hex(n int) string {
	base := "0123456789abcdef"
	if n < 16 {
		return base[n : n+1]
	}
	return Hex(n/16) + base[n%16:n%16+1]
}
