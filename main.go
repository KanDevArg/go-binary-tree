package main

import (
	"fmt"
	"sync"
)

type node struct {
	value int
	right *node
	left  *node
}

func findValue(joint *node, value int, ch chan<- int, ca chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	abort:=<-ca

	if abort {
		fmt.Println("Should abort!!")
		return
	}

	found := false

	if (*joint).value == value {
		ch <- value
		found = false
	}



	if (*joint).left == nil && (*joint).right == nil {
			return
	} else {
		if (*joint).right != nil {
			wg.Add(1)
			go func() {
				ca <- found
			}()
			go findValue((*joint).right, value, ch, ca, wg)
		}

		if (*joint).left != nil {
			wg.Add(1)
			go func() {
				ca <- found
			}()
			go findValue((*joint).left, value, ch, ca, wg)
		}
	}
}

func main(){
	tree1 := &node{value: 10, right: &node{value: 11, right: nil, left: nil}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: &node{value: 11, right: &node{value: 11, right: &node{value: 10, right: &node{value: 11, right: nil, left: nil}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: nil}}}, left: &node{value: 10, right: &node{value: 11, right: nil, left: nil}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: &node{value: 11, right: &node{value: 11, right: &node{value: 10, right: &node{value: 11, right: nil, left: nil}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: nil}}}, left: nil}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: nil}}}}}}}, left: &node{value: 11, right: nil, left: &node{ value:10, right: nil, left: nil}}}}}}
	ch := make(chan int)
	ca := make(chan bool)
	var wg sync.WaitGroup

	go func () {
		ca<-false
	}()

	wg.Add(1)
	go findValue(tree1,11, ch, ca, &wg )

	go func () {
		defer close(ch)
		defer close(ca)
		defer wg.Wait()
	}()

	for v := range ch {
		fmt.Println("value", v)
	}
}