package main

import (
	"fmt"
	"gofmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	for true {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		t := time.Now()
		if t.Hour() <= 9 {
			fmt.Printf("0")
		}
		fmt.Printf("%d:", t.Hour())
		if t.Minute() <= 9 {
			fmt.Printf("0")
		}
		fmt.Printf("%d:", t.Minute())
		if t.Second() <= 9 {
			fmt.Printf("0")
		}
		fmt.Printf("%d\n", t.Second())
	}
}
