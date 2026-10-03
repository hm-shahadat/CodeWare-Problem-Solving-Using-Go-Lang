// Problem name : Write a function that removes the spaces from the string, then return the resultant string.

// Examples (Input -> Output):

// "8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
// "8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
// "8aaaaa dddd r     " -> "8aaaaaddddr"

//  link: https://www.codewars.com/kata/57eae20f5500ad98e50002c5/train/go

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
