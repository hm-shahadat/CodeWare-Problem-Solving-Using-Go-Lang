package kata

import "fmt"

func NoSpace(word string) string {

	var str []rune
	for i, c := range word {
		if word[i] != ' ' {

			str = append(str, c)
		}
	}
	return string(str)
}

func main() {
	mystr := NoSpace("shahadat hossain gazi")

	fmt.Println(mystr)
}
