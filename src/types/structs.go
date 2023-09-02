package main
import "fmt"

type TextFile struct {
	name string
	size float64
	characters int
	words int
	lines int
}

func main() {
	file := TextFile{"structs.go", 14496, 174, 63, 14}
	fmt.Println(file)
}