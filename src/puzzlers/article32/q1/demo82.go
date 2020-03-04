package main

import (
	"fmt"
	"strings"
)

func main() {
	comment := "Package io provides basic interfaces to I/O primitives. " +
		"Its primary job is to wrap existing implementations of such primitives, " +
		"such as those in package os, " +
		"into shared public interfaces that abstract the functionality, " +
		"plus some other related primitives."

	// 示例1
	fmt.Println("New a string reader and name it \"reader1\" ...")
	reader1 := strings.NewReader(comment)
	buf1 := make([]byte,7)
	n,err := reader1.Read(buf1)
	var offset1,index1 int64
	executeIfNoErr(err,func()){

	}
}

func executeIfNoErr(err error,f func())  {
	if err!= nil{
		fmt.Printf("error: %v\n",err)
		return
	}
	f()
}
