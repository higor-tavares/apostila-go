package main
import (
	"fmt"
	"io"
)

type ASCIIReader struct {}

func (r ASCIIReader) Read(p []byte) (int, error) {
	return len(p), nil
}
func readString(r io.Reader) string {
	p := make([]byte, 5)
	p[0] = 'H'
	p[1] = 'I'
	p[2] = 'G'
	p[3] = 'O'
	p[4] = 'R'
	size, _ := r.Read(p)
	fmt.Println("You read a string with length ", size)
	return string(p)
}

func main() {
	reader := ASCIIReader{}
	fmt.Println(readString(reader))
}