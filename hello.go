package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello!")
}
func main() {
	go sayHello()
	time.Sleep(5 * time.Second)
	go lilgoroutine()
	time.Sleep((5 * time.Second))
}

func lilgoroutine() {
	fmt.Println(("goroutine numbver2"))
}
