package main

import (
	"fmt"
	"gofmt"
	"time"
)

func main() {
    for true {
    	t := time.Now()
    	fmt.Printf("%d:", t.Hour())
    	fmt.Printf("%d:", t.Minute())
    	fmt.Printf("%d\n", t.Second())
    }
}
