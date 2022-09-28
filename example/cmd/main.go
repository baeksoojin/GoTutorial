package main

import (
	tutorial "example/pkg/tutorial"
	"sync"
)

func main() {

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
	tutorial.Say("world")

	wg.Wait() //wg.Done으로 WaitGroup 객체가 모두 종료될때까지 기다린다.

	//// ex
	//v := tutorial.Vertex{3, 4}
	//fmt.Println(v.Abs())
	//
	//// add function
	//v.Add()
	//fmt.Println(v)
	//fmt.Println(v.Abs())
	////subtraction function
	//v.Sub()
	//fmt.Println(v)
	//fmt.Println(v.Abs())

}
