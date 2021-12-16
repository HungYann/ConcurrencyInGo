package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	for _, salution := range []string{"hello","greetings","good day"}{
		wg.Add(1)
		go func(salution string) {
			defer wg.Done()
			fmt.Println(salution)
		}(salution)
	}
	wg.Wait()
}
