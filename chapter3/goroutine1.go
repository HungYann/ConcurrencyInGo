package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	for _, salution := range []string{"hello","greetins","good day"}{
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salution)
		}()
	}
	wg.Wait()
}
