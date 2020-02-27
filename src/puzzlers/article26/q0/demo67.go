package main

import "fmt"

func main() {
	coordinateWithWaitGroup()
}

func coordinateWithWaitGroup() {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup\n]", num)

}
