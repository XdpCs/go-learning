package main

import "fmt"

type Fyy struct {
	Bag
}

type Bag interface {
	Hello()
}

type Xdp struct {
	Fyy
}

func main() {
	var bag Bag = new(Fyy)
	fmt.Println(bag)
	bag.Hello()
}
