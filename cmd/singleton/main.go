package main

import (
	"fmt"
	"sync"

	"github.com/domhoward14/creational_design_patterns/internal/singleton/db"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go execQuery(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines have finished executing.")
}

func execQuery(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := db.GetSingleton()
	res := s.Query(fmt.Sprintf("SELECT * FROM test_table_%d;", i))
	fmt.Println(res)
}
