package main

import (
	"fmt"
	"sync"
)

func main(){
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decerment := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	//增量
	var arithmetic sync.WaitGroup
	for i:=0;i<=5;i++{
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	//减量
	for i:=0;i<=5;i++{
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decerment()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}