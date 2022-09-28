package tutorial

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Printer1(val int) {
	for i := 0; i < 2; i++ {
		fmt.Println(val)
	}
}

func Printer2(val int) {
	for i := 0; i < 2; i++ {
		fmt.Println(val)
	}
}
