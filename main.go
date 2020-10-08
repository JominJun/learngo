package main

import (
	"github.com/JominJun/learngo/accounts"
	"fmt"
	//"log"
)

func main(){
	account := accounts.NewAccount("minjun")
	account.Deposit(10)
	fmt.Println(account)
}