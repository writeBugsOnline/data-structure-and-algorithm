package main

import (
	"fmt"
	"strings"
	"time"
)


var c =0
var num =0

type stu struct {
	Name string
	Age int
}
func main(){

}

func test1()  {
	chars :="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cslice :=strings.Split(chars,"")
	syn1 :=make(chan int,1)
	syn2 :=make(chan int,1)
	syn2<-1

	go print_num(syn1,syn2)
	go print_char(cslice,syn1,syn2)

	time.Sleep(10)

	stus := []stu{
		{
			Name: "lisi",
			Age:24,
		},
		{
			Name:"zhangsan",
			Age:15,
		},
	}
	m := make(map[string]*stu)

	for i:=0;i<len(stus);i++ {
		m[stus[i].Name] = &stus[i]
	}

	for i :=range stus{
		m[stus[i].Name] = &stus[i]
	}
	//
	// for _,stu :=range stus{
	// 	stu.Age=stu.Age+10
	// }

}

func print_char(cslice []string, syn1 chan int, syn2 chan int) {
	for c <26 {
		<-syn1
		fmt.Print( cslice[c])
		c++
		fmt.Print( cslice[c])
		c++
		syn2 <-1
	}

}

func print_num(syn1 chan int, syn2 chan int) {
	for num <=26 {
		<-syn2
		num++
		fmt.Print(num)
		num++
		fmt.Print(num)
		syn1 <- 1
	}



}

