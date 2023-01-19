package controller

import "fmt"

type Controller struct{}

type Interface interface {
	Controllerfunc()
}

func (c *Controller) Controllerfunc() {
	fmt.Println("this is controller function")
}

func NewController() Interface {
	return &Controller{}
}