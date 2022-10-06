package main

import (
	tutorial "example/pkg/tutorial"
	"fmt"
	"github.com/spf13/pflag"
	"sync"
)

func main() {

	//flag ex
	flags := pflag.NewFlagSet("", pflag.ExitOnError)

	// goroutine ex
	wg := sync.WaitGroup{} // WaitGroup 생성
	go func() {
		wg.Add(1) //waitgroup이 가진 값을 arg만큼 증가
		defer wg.Done()
		tutorial.Printer1(1)
	}()
	go func() {
		wg.Add(1) //waitgroup이 가진 값을 arg만큼 증가
		defer wg.Done()
		tutorial.Printer2(2) //실행되지 않음
	}()
	go func() {
		wg.Add(1) //waitgroup이 가진 값을 arg만큼 증가
		defer wg.Done()
		tutorial.Say("hello")
	}()
	go func() { //test flags
		wg.Add(1) //waitgroup이 가진 값을 arg만큼 증가
		defer wg.Done()
		t := tutorial.NewSet(flags)
		fmt.Println("get value", t.GetValue())
	}()
	go func() { //test flags -> 2개 test
		wg.Add(1) //waitgroup이 가진 값을 arg만큼 증가
		defer wg.Done()
		t2 := tutorial.NewSet2(flags)
		fmt.Println("get value2", t2.GetValue())
	}()
	tutorial.Say("skjfjdfjkdjfjkjkjkjjk")

	wg.Wait() //wg.Done으로 WaitGroup 객체가 모두 종료될때까지 기다린다.

}
