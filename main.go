package main

import (
	MyArray "LearnAlgorithm/Array"
	MyString "LearnAlgorithm/String"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	var words = []string{"cat", "bt", "hat", "tree"}
	var chars string = "atach"
	//var len = countCharacters(words, chars)
	var len = MyArray.NotExhaustiveAttackMethod(words, chars)
	fmt.Println(len)
	var res = MyString.IsValid("sadfsdf")
	fmt.Println(res)
}
