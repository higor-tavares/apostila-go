package main
import (
	"fmt"
	"regexp"
)

func escape(anyText string) string {
	expression := regexp.MustCompile("[a-zA-Zàáâãèéêîíìçôòóûùú]")
	result := expression.ReplaceAllStringFunc(
		anyText,
		func(s string) string {
			switch (s) {
				case "à","á","â","ã":
					return "a"
				case "è","é","ê":
					return "e"
				case "í","î","ì":
					return "i"
				case "ô","ò","ó","õ":
					return "o"
				case "û","ù","ú":
					return "u"
				case "ç":
					return "c"
				default:
					return s

			}
		})
	return result
}

func main() {
	n1 := "assunção"
	n2 := "pelé"
	fmt.Println(escape(n1))
	fmt.Println(escape(n2))
}