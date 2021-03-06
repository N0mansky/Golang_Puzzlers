package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	defer func() {
		fmt.Println("Enter defer function.")

		// recover 函数的正确用法
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()

	// recover 函数的错误用法
	fmt.Printf("no panic: %v\n", recover())

	// 引发 panic
	panic(errors.New("something wrong"))
}
