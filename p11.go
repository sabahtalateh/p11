package main

import (
	"github.com/sabahtalateh/p1/p11/p12"
	"github.com/sabahtalateh/p1/p11/p12/p13"
)

type P11 struct {
}

type P111 struct {
}

func main() {
	p := p12.P12{}
	println(&p)

	pp := p13.P13{}
	println(&pp)
}
