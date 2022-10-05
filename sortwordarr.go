package piscine

// import "fmt"

func SortWordArr(a []string) {
	done := false
	for !done {
		done = true
		for i := 0; i < len(a)-1; i++ {
			if a[i+1] < a[i] {
				a[i+1], a[i] = a[i], a[i+1]
				done = false
			}
		}
	}
}
