package main

import (
	"github.com/JominJun/learngo/mydict"
	"fmt"
)

func main(){
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")

	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}