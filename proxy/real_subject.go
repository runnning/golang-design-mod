package proxy

import "fmt"

type RealSubject struct {
}

func (rs *RealSubject) Request() {
	fmt.Println("RealSubject: Handling request")
}
