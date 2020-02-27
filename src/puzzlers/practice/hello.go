package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//timeStr := time.Now().Format("20060102150405")
	for i := 0; i <= 30; i++ {
		s2 := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		fmt.Println(s2)

	}
	//fmt.Println(timeStr)
}
